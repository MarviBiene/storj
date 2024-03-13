pipeline {
    parameters {
        string(name: 'BUILD_TYPE', defaultValue: '', description: 'Group of tests which should be executed. Use pre-merge, VALIDATE or empty.')
    }
    agent {
        docker {
            label 'main'
            image 'storjlabs/ci:latest'
            alwaysPull true
            args '-u root:root --cap-add SYS_PTRACE -v "/tmp/gomod":/go/pkg/mod -v "/tmp/npm":/npm --tmpfs "/tmp:exec,mode=777"'
        }
    }
    options {
        timeout(time: 4, unit: 'HOURS')
        skipDefaultCheckout(true)
    }
    environment {
        BUILD_TYPE = "${params.BUILD_TYPE}"
        NPM_CONFIG_CACHE = '/npm/cache'
        GOTRACEBACK = 'all'
        COCKROACH_MEMPROF_INTERVAL=0
    }
    stages {
        stage('Checkout') {
            steps {
                // Delete any content left over from a previous run.
                sh "chmod -R 777 ."
                // Bash requires extglob option to support !(.git) syntax,
                // and we don't want to delete .git to have faster clones.
                sh 'bash -O extglob -c "rm -rf !(.git)"'

                checkout scm

                sh 'mkdir -p .build'

                // download dependencies
                sh 'go mod download'

                // pre-check that we cannot do at a later stage reliably
                sh 'check-large-files'
            }
        }
        stage('Gerrit status') {
            steps {
                withCredentials([sshUserPrivateKey(credentialsId: 'gerrit-trigger-ssh', keyFileVariable: 'SSH_KEY', usernameVariable: 'SSH_USER')]) {
                    sh './scripts/gerrit-status.sh $BUILD_TYPE start 0'
                }
            }
        }
        stage('Build Web') {
            when {
                anyOf {
                    equals expected: "pre-merge", actual: params.BUILD_TYPE
                    equals expected: "", actual: params.BUILD_TYPE
                }
            }
            // The build code depends on the following assets being loaded.
            parallel {
                stage('web/satellite') {
                    steps {
                        sh './web/satellite/build.sh'
                    }
                }

                stage('wasm') {
                    steps {
                        sh './testsuite/wasm/start.sh'
                    }
                }

                stage('web/storagenode') {
                    steps {
                        sh './web/storagenode/build.sh'
                    }
                }

                stage('web/multinode') {
                    steps {
                        sh './web/multinode/build.sh'
                    }
                }

                stage('satellite/admin/ui') {
                    steps {
                        sh './satellite/admin/ui/build.sh'
                    }
                }
            }
        }

       stage('Build') {
            when {
                anyOf {
                    equals expected: "verify", actual: params.BUILD_TYPE
                    equals expected: "", actual: params.BUILD_TYPE
                }
            }
            parallel {
                stage('go') {
                    steps {
                        // use go test to build all the packages, including tests
                        sh 'go test -v -p 16 -bench XYZXYZXYZXYZ -run XYZXYZXYZXYZ ./...'
                    }
                }
                stage('go -race') {
                    steps {
                        // use go test to build all the packages, including tests
                        sh 'go test -v -p 16 -bench XYZXYZXYZXYZ -run XYZXYZXYZXYZ -race ./...'

                        // install storj-sim
                        sh 'go install -race -v storj.io/storj/cmd/satellite '+
                                'storj.io/storj/cmd/storagenode ' +
                                'storj.io/storj/cmd/storj-sim ' +
                                'storj.io/storj/cmd/versioncontrol ' +
                                'storj.io/storj/cmd/uplink ' +
                                'storj.io/storj/cmd/identity ' +
                                'storj.io/storj/cmd/certificates ' +
                                'storj.io/storj/cmd/multinode'
                    }
                }
                stage('go -race gateway') {
                    steps {
                        // install gateway for storj-sim
                        sh 'go install -race -v storj.io/gateway@latest'
                    }
                }

                stage('db') {
                    steps {
                        dir('.build') {
                            sh 'cockroach start-single-node --insecure --store=type=mem,size=4GiB --listen-addr=localhost:26256 --http-addr=localhost:8086 --cache 1024MiB --max-sql-memory 1024MiB --background'
                            sh 'cockroach start-single-node --insecure --store=type=mem,size=4GiB --listen-addr=localhost:26257 --http-addr=localhost:8087 --cache 1024MiB --max-sql-memory 1024MiB --background'
                        }
                    }
                }
            }
        }

                stage('Lint') {
                    when {
                        anyOf {
                            equals expected: "verify", actual: params.BUILD_TYPE
                            equals expected: "", actual: params.BUILD_TYPE
                        }
                    }
                    steps {
                        sh 'check-mod-tidy'
                        sh 'check-copyright'
                        sh 'check-imports -race ./...'
                        sh 'check-peer-constraints -race'
                        sh 'check-atomic-align ./...'
                        sh 'check-monkit ./...'
                        sh 'check-errs ./...'
                        sh 'staticcheck ./...'
                        sh 'golangci-lint --config /go/ci/.golangci.yml -j=2 run'
                        sh 'check-downgrades'
                        sh 'make check-monitoring'
                        sh 'make test-wasm-size'

                        sh 'protolock status'
                        sh './scripts/check-package-lock.sh'

                        // go-licenses by default has AGPL3 in the forbidden list, hence we need to explicitly allow `storj.io/storj`.
                        sh 'go-licenses check --ignore "storj.io/storj" ./...'
                    }
                }

                stage('Satellite UI Tests') {
                    when {
                        anyOf {
                            anyOf {
                                equals expected: "verify", actual: params.BUILD_TYPE
                                equals expected: "", actual: params.BUILD_TYPE
                            }
                            anyOf {
                                branch 'main'
                                branch pattern: "release-.*", comparator: "REGEXP"
                                changeset "testsuite/playwright-ui/**"
                                changeset "web/**"
                                changeset "satellite/console/**"
                            }
                        }
                    }
                    environment {
                        STORJ_TEST_COCKROACH = 'cockroach://root@localhost:26256/uitestcockroach?sslmode=disable'
                        STORJ_TEST_COCKROACH_NODROP = 'true'
                        STORJ_TEST_POSTGRES = 'omit'
                        STORJ_TEST_LOG_LEVEL = 'debug'
                    }

                    steps {
                        sh 'cockroach sql --insecure --host=localhost:26256 -e \'create database uitestcockroach;\''
                        sh 'make test-satellite-ui'
                    }
                }

                stage('Cross Compile') {
                    when {
                        anyOf {
                            equals expected: "pre-merge", actual: params.BUILD_TYPE
                            equals expected: "", actual: params.BUILD_TYPE
                        }
                    }
                    steps {
                        // verify most of the commands, we cannot check everything since some of them
                        // have a C dependency and we don't have cross-compilation in storj/ci image
                        sh 'check-cross-compile storj.io/storj/cmd/uplink storj.io/storj/cmd/satellite storj.io/storj/cmd/storagenode-updater storj.io/storj/cmd/storj-sim'
                    }
                }
                stage('Tests') {
                    when {
                        anyOf {
                            equals expected: "verify", actual: params.BUILD_TYPE
                            equals expected: "", actual: params.BUILD_TYPE
                        }
                    }
                    environment {
                        STORJ_TEST_HOST = '127.0.0.20;127.0.0.21;127.0.0.22;127.0.0.23;127.0.0.24;127.0.0.25'
                        STORJ_TEST_COCKROACH = 'cockroach://root@loc0alhost:26256/testcockroach?sslmode=disable;' +
                            'cockroach://root@localhost:26257/testcockroach?sslmode=disable;'
                        STORJ_TEST_COCKROACH_NODROP = 'true'
                        STORJ_TEST_COCKROACH_ALT = 'cockroach://root@localhost:26260/testcockroach?sslmode=disable'
                        STORJ_TEST_POSTGRES = 'postgres://postgres@localhost/teststorj?sslmode=disable'
                        STORJ_TEST_LOG_LEVEL = 'info'
                        COVERFLAGS = "${ env.BRANCH_NAME == 'main' ? '-coverprofile=.build/coverprofile -coverpkg=storj.io/storj/private/...,storj.io/storj/satellite/...,storj.io/storj/storagenode/...,storj.io/storj/versioncontrol/...' : ''}"
                        GOEXPERIMENT = 'nocoverageredesign'
                    }
                    steps {
                        sh 'cockroach sql --insecure --host=localhost:26256 -e \'create database testcockroach;\''
                        sh 'cockroach sql --insecure --host=localhost:26257 -e \'create database testcockroach;\''

                        sh 'psql -U postgres -c \'create database teststorj;\''

                        sh 'use-ports -from 1024 -to 10000 &'

                        sh 'go test -parallel 4 -p 6 -vet=off $COVERFLAGS -timeout 32m -json -race ./... 2>&1 | tee .build/tests.json | xunit -out .build/tests.xml'
                    }

                    post {
                        always {
                            archiveArtifacts artifacts: '.build/tests.json'
                            sh script: 'cat .build/tests.json | tparse -all -slow 100', returnStatus: true
                            junit '.build/tests.xml'

                            script {
                                if(fileExists(".build/coverprofile")){
                                    sh script: 'filter-cover-profile < .build/coverprofile > .build/clean.coverprofile', returnStatus: true
                                    sh script: 'gocov convert .build/clean.coverprofile > .build/cover.json', returnStatus: true
                                    sh script: 'gocov-xml  < .build/cover.json > .build/cobertura.xml', returnStatus: true
                                    sh script: 'gzip .build/clean.coverprofile -c > .build/clean.coverprofile.gz'
                                    archiveArtifacts artifacts: '.build/clean.coverprofile.gz'
                                    cobertura coberturaReportFile: '.build/cobertura.xml',
                                        lineCoverageTargets: '70, 60, 50',
                                        autoUpdateHealth: false,
                                        autoUpdateStability: false,
                                        failUnhealthy: true
                                }
                            }
                        }
                    }
                }

                stage('Check Benchmark') {
                    when {
                        anyOf {
                            equals expected: "pre-merge", actual: params.BUILD_TYPE
                            equals expected: "", actual: params.BUILD_TYPE
                        }
                    }
                    environment {
                        STORJ_TEST_COCKROACH = 'omit'
                        STORJ_TEST_POSTGRES = 'postgres://postgres@localhost/benchstorj?sslmode=disable'
                    }
                    steps {
                        sh 'psql -U postgres -c \'create database benchstorj;\''
                        sh 'go test -parallel 1 -p 1 -vet=off -timeout 20m -short -run XYZXYZXYZXYZ -bench . -benchtime 1x ./...'
                    }
                }

                stage('Integration') {
                    when {
                        anyOf {
                            equals expected: "pre-merge", actual: params.BUILD_TYPE
                            equals expected: "", actual: params.BUILD_TYPE
                        }
                    }
                    environment {
                        // use different hostname to avoid port conflicts
                        STORJ_NETWORK_HOST4 = '127.0.0.2'
                        STORJ_NETWORK_HOST6 = '127.0.0.2'

                        STORJ_SIM_POSTGRES = 'postgres://postgres@localhost/teststorj2?sslmode=disable'
                    }

                    steps {
                        sh 'psql -U postgres -c \'create database teststorj2;\''
                        sh 'make test-sim'

                        // sh 'make test-certificates' // flaky
                    }
                }

                stage('Cockroach Integration') {
                    when {
                        anyOf {
                            equals expected: "", actual: params.BUILD_TYPE
                        }
                    }
                    environment {
                        STORJ_NETWORK_HOST4 = '127.0.0.4'
                        STORJ_NETWORK_HOST6 = '127.0.0.4'

                        STORJ_SIM_POSTGRES = 'cockroach://root@localhost:26257/testcockroach4?sslmode=disable'
                    }

                    steps {
                        sh 'cockroach sql --insecure --host=localhost:26257 -e \'create database testcockroach4;\''
                        sh 'make test-sim'
                        sh 'cockroach sql --insecure --host=localhost:26257 -e \'drop database testcockroach4;\''
                    }
                }

                stage('Integration Redis unavailability') {
                    when {
                        anyOf {
                            equals expected: "pre-merge", actual: params.BUILD_TYPE
                            equals expected: "", actual: params.BUILD_TYPE
                        }
                    }
                    environment {
                        // use different hostname to avoid port conflicts
                        STORJ_NETWORK_HOST4 = '127.0.0.6'
                        STORJ_NETWORK_HOST6 = '127.0.0.6'
                        STORJ_REDIS_PORT = '7379'

                        STORJ_SIM_POSTGRES = 'postgres://postgres@localhost/teststorj6?sslmode=disable'
                    }

                    steps {
                        sh 'psql -U postgres -c \'create database teststorj6;\''
                        sh 'make test-sim-redis-unavailability'
                    }
                }

                stage('Backwards Compatibility') {
                    when {
                        anyOf {
                            equals expected: "pre-merge", actual: params.BUILD_TYPE
                            equals expected: "", actual: params.BUILD_TYPE
                        }
                    }
                    environment {
                        STORJ_NETWORK_HOST4 = '127.0.0.3'
                        STORJ_NETWORK_HOST6 = '127.0.0.3'

                        STORJ_SIM_POSTGRES = 'postgres://postgres@localhost/teststorj3?sslmode=disable'
                        STORJ_MIGRATION_DB = 'postgres://postgres@localhost/teststorj3?sslmode=disable&options=--search_path=satellite/0/meta'
                    }

                    steps {
                        sh 'psql -U postgres -c \'create database teststorj3;\''
                        sh 'make test-sim-backwards-compatible'
                    }
                }

                stage('Cockroach Backwards Compatibility') {
                    when {
                        anyOf {
                            equals expected: "", actual: params.BUILD_TYPE
                        }
                    }
                    environment {
                        STORJ_NETWORK_HOST4 = '127.0.0.5'
                        STORJ_NETWORK_HOST6 = '127.0.0.5'

                        STORJ_SIM_POSTGRES = 'cockroach://root@localhost:26257/testcockroach5?sslmode=disable'
                        STORJ_MIGRATION_DB = 'postgres://root@localhost:26257/testcockroach5/satellite/0/meta?sslmode=disable'
                    }

                    steps {
                        sh 'cockroach sql --insecure --host=localhost:26257 -e \'create database testcockroach5;\''
                        sh 'make test-sim-backwards-compatible'
                        sh 'cockroach sql --insecure --host=localhost:26257 -e \'drop database testcockroach5;\''
                    }
                }
        stage('Post') {
            parallel {
                stage('Lint') {
                    steps {
                        sh 'check-clean-directory'
                    }
                }
            }
        }
    }
    post {
        success {
            withCredentials([sshUserPrivateKey(credentialsId: 'gerrit-trigger-ssh', keyFileVariable: 'SSH_KEY', usernameVariable: 'SSH_USER')]) {
                sh './scripts/gerrit-status.sh $BUILD_TYPE success +2'
            }
        }
        failure {
            withCredentials([sshUserPrivateKey(credentialsId: 'gerrit-trigger-ssh', keyFileVariable: 'SSH_KEY', usernameVariable: 'SSH_USER')]) {
                sh './scripts/gerrit-status.sh $BUILD_TYPE failure -2'
            }
        }
        aborted {
            withCredentials([sshUserPrivateKey(credentialsId: 'gerrit-trigger-ssh', keyFileVariable: 'SSH_KEY', usernameVariable: 'SSH_USER')]) {
                sh './scripts/gerrit-status.sh $BUILD_TYPE failure -2'
            }
        }
    }
}
