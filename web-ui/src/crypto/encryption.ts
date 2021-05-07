import { Encoding } from "./encoding"

export interface IEncryptedPayload {
    initializationVector: string
    encrypted: string
}

export class Encryption {
    private static readonly importExportFormat = 'raw'
    private static readonly algo = 'AES-GCM'
    private static readonly bits = 128
    private static readonly initializationVectorLength = 96

    private static readonly keyUsages: KeyUsage[] = [ 'encrypt', 'decrypt' ]

    public static async randomKey(): Promise<CryptoKey> {
        const params: AesKeyGenParams = { 
            name: this.algo,
            length: this.bits,
        }
        return await crypto.subtle.generateKey(params, true, this.keyUsages)
    }

    public static async exportKey(key: CryptoKey): Promise<ArrayBuffer> {
        return await crypto.subtle.exportKey(this.importExportFormat, key)
    }

    public static async importKey(buffer: ArrayBuffer): Promise<CryptoKey> {
        const algo: AesKeyAlgorithm = {
            name: this.algo,
            length: this.bits,
        }

        return await crypto.subtle.importKey(this.importExportFormat, buffer, algo, true, this.keyUsages)
    }

    public static async encrypt(plaintext: string, key: CryptoKey): Promise<IEncryptedPayload> {
        const initVect = window.crypto.getRandomValues(new Uint8Array(this.initializationVectorLength / 8));

        const algo: AesGcmParams = { 
            name: this.algo,
            iv: initVect,
        }

        const enc = new TextEncoder()
        const plaintextBytes = enc.encode(plaintext)
       
        const encrypted = await crypto.subtle.encrypt(algo, key, plaintextBytes)

        return { 
            initializationVector: Encoding.encode(initVect),
            encrypted: Encoding.encode(encrypted),
        }
    } 

    public static async decrypt(encryptedData: IEncryptedPayload, key: CryptoKey): Promise<string> {
        const algo: AesGcmParams = { 
            name: this.algo,
            iv: Encoding.decode(encryptedData.initializationVector),
        }

        const plaintextBytes = await crypto.subtle.decrypt(algo, key, Encoding.decode(encryptedData.encrypted))

        const dec = new TextDecoder()
        const plaintext = dec.decode(plaintextBytes)

        return plaintext
    }

    
}