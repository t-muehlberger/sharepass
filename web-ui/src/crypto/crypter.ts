import CryptoJS from "crypto-js"

export class Crypter {
    public static randomKey(): string {
        return "supersecret"
    }

    public static encrypt(plaintext: string, key: string): string {
        let cryptoParams = CryptoJS.AES.encrypt(plaintext, key)
        return cryptoParams.toString()
    } 

    public static decrypt(cipher: string, key: string): string {
        let bytes = CryptoJS.AES.decrypt(cipher, key)
        return bytes.toString(CryptoJS.enc.Utf8);
    }
}