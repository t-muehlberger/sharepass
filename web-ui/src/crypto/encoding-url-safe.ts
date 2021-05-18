import { Encoding } from "./encoding";

export class EncodingUrlSafe {
    // https://www.base64url.com
    public static encode(data: ArrayBuffer): string {
        const base64 = Encoding.encode(data)
        const urlSafe = base64
            .replace(/\+/g, '-')
            .replace(/\//g, '_')
            .replace(/=+$/g, '');
        return urlSafe
    }

    public static decode(data: string): ArrayBuffer {
        const base64 = data
            .replace(/-/g, '+')
            .replace(/_/g, '/')

        return Encoding.decode(base64)
    }
}
