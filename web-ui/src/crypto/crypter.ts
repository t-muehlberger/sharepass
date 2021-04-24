import AES from "crypto-js/aes"
import UTF8Encoder from "crypto-js/enc-utf8"

export class Crypter {
    public static randomKey(): string {
        const bytes = new Uint8Array(32)
        window.crypto.getRandomValues(bytes)
        return this.numbersToHex(bytes)
    }

    public static numbersToHex(data: Uint8Array): string {
        const str = [...data]
            .map((b: number):string => b.toString(16).padStart(2, "0"))
            .join("");
        return str;
    }

    public static encrypt(plaintext: string, key: string): string {
        const cryptoParams = AES.encrypt(plaintext, key)
        return cryptoParams.toString()
    } 

    public static decrypt(cipher: string, key: string): string {
        const bytes = AES.decrypt(cipher, key)
        const decrypted = bytes.toString(UTF8Encoder);
        return decrypted
    }
}