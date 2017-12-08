# cfgcrypt

CfgCrypt or config crypt is a cli tool to encrypt values in a text configuration file for use within a secure application.

## Concept

 Write your configuration file in whatever text format you prefer, then wrap any values that you would like to keep secret in prefix and postfix delimiters that occur no where else in your file. Then you run cfgcrypt with your delimiters on the file to encrypt the variables you want hidden. From there your application decodes the secret values using the configuration file and a key file.

 ## Decryption logic

 Pseudo-code example of decryption process for CBC encryption:

 ```
 configData = io.ReadConfigFile(fileName)
 key = io.ReadFileBytes(fileName + ".key")
 decoded = base64.decode(configData.secretValue)
 iv = getFirst16Bytes(decoded)
 encrypted = getBytesAfterFirst16(decoded)
 cipher = encryption.Mode("AES/CBC/PKCS7Padding")
 configData.secretValue = cipher.Decrypt(iv, encrypted, key)
 ```

 ## Usage

To use boolean parameters set -boolvar=true/false
 ```
 cfgcrypt [textfile] ...
    textfile    Text file to encrypt. (required)
   -debug bool
     	Display detailed error messages
   -force bool
     	Overwrite key file if found
   -key string
     	Base64 encoded encryption key, if not specified one will be generated
   -postfix string
     	Post string denoting end of value to be encrypted (default "}}#")
   -prefix string
     	Prefix string denoting start of value to be encrypted (default "#{{")
```

## Example

Examples are located in the `examples` folder with the original unencrypted files in `examples/original` and the encrypted output in `examples/encrypted`

## How To Build

To build run
```bash
godep save
godep go build
```

## Future development

I'm considering the following upgrades:
-   Support for more encryption modes/algorithms
