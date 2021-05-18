export class Encoding {
    public static encode(key: ArrayBuffer): string {
        const str = String.fromCharCode(...new Uint8Array(key))
        const base64 = btoa(str)
        return base64
    }

    public static decode(key: string): ArrayBuffer {
        const origStr = atob(key)

        const buffer = new Uint8Array(origStr.length) 
        for (let i = 0; i < buffer.length; i++) {
            buffer[i] = origStr.charCodeAt(i)
        } 

        return buffer
    }
}
