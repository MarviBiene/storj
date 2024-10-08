// Copyright (C) 2024 Storj Labs, Inc.
// See LICENSE for copying information.

import { ObjectLockMode } from '@aws-sdk/client-s3';

export class Retention {
    mode: ObjectLockMode | '';
    retainUntil: Date;

    constructor(mode: ObjectLockMode | '', retainUntil: Date) {
        this.mode = mode;
        this.retainUntil = retainUntil;
    }

    public static empty(): Retention {
        return new Retention('', new Date());
    }

    // returns whether the retention configuration is enabled.
    public enabled(): boolean {
        return this.mode === ObjectLockMode.COMPLIANCE || this.mode === ObjectLockMode.GOVERNANCE;
    }

    // returns whether the retention configuration is enabled
    // and active as of the current time.
    public active(): boolean {
        return this.enabled() && new Date() < this.retainUntil;
    }
}