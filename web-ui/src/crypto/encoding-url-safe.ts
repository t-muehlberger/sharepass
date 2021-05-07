import { Encoding } from "./encoding";

export class EncodingUrlSafe {
    public static encode(key: ArrayBuffer): string {
        const base64 = Encoding.encode(key)
        const urlSafe = base64
            .replace(/\+/g, '-')
            .replace(/\//g, '.')
            .replace(/=+$/g, '');
        return urlSafe
    }

    public static decode(key: string): ArrayBuffer {
        const base64 = key
            .replace(/-/g, '+')
            .replace(/\./g, '/')

        return Encoding.decode(base64)
    }
}
