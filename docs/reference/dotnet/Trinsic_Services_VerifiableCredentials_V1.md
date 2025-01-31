#### [Trinsic](index.md 'index')
## Trinsic.Services.VerifiableCredentials.V1 Namespace
### Classes
<a name='Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest'></a>
## CheckStatusRequest Class
request object to update the status of the revocation entry  
```csharp
public sealed class CheckStatusRequest :
Google.Protobuf.IMessage<Trinsic.Services.VerifiableCredentials.V1.CheckStatusRequest>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.VerifiableCredentials.V1.CheckStatusRequest>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.VerifiableCredentials.V1.CheckStatusRequest>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; CheckStatusRequest  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[CheckStatusRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest 'Trinsic.Services.VerifiableCredentials.V1.CheckStatusRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[CheckStatusRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest 'Trinsic.Services.VerifiableCredentials.V1.CheckStatusRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[CheckStatusRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest 'Trinsic.Services.VerifiableCredentials.V1.CheckStatusRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest_CredentialStatusIdFieldNumber'></a>
## CheckStatusRequest.CredentialStatusIdFieldNumber Field
Field number for the "credential_status_id" field.
```csharp
public const int CredentialStatusIdFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
### Properties
<a name='Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest_CredentialStatusId'></a>
## CheckStatusRequest.CredentialStatusId Property
the credential status id  
```csharp
public string CredentialStatusId { get; set; }
```
#### Property Value
[System.String](https://docs.microsoft.com/en-us/dotnet/api/System.String 'System.String')
  
  
<a name='Trinsic_Services_VerifiableCredentials_V1_CheckStatusResponse'></a>
## CheckStatusResponse Class
response object for update of status of revocation entry  
```csharp
public sealed class CheckStatusResponse :
Google.Protobuf.IMessage<Trinsic.Services.VerifiableCredentials.V1.CheckStatusResponse>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.VerifiableCredentials.V1.CheckStatusResponse>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.VerifiableCredentials.V1.CheckStatusResponse>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; CheckStatusResponse  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[CheckStatusResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CheckStatusResponse 'Trinsic.Services.VerifiableCredentials.V1.CheckStatusResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[CheckStatusResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CheckStatusResponse 'Trinsic.Services.VerifiableCredentials.V1.CheckStatusResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[CheckStatusResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CheckStatusResponse 'Trinsic.Services.VerifiableCredentials.V1.CheckStatusResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_VerifiableCredentials_V1_CheckStatusResponse_RevokedFieldNumber'></a>
## CheckStatusResponse.RevokedFieldNumber Field
Field number for the "revoked" field.
```csharp
public const int RevokedFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
### Properties
<a name='Trinsic_Services_VerifiableCredentials_V1_CheckStatusResponse_Revoked'></a>
## CheckStatusResponse.Revoked Property
indicates if the status is revoked  
```csharp
public bool Revoked { get; set; }
```
#### Property Value
[System.Boolean](https://docs.microsoft.com/en-us/dotnet/api/System.Boolean 'System.Boolean')
  
  
<a name='Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest'></a>
## CreateProofRequest Class
Create Proof  
```csharp
public sealed class CreateProofRequest :
Google.Protobuf.IMessage<Trinsic.Services.VerifiableCredentials.V1.CreateProofRequest>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.VerifiableCredentials.V1.CreateProofRequest>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.VerifiableCredentials.V1.CreateProofRequest>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; CreateProofRequest  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[CreateProofRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest 'Trinsic.Services.VerifiableCredentials.V1.CreateProofRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[CreateProofRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest 'Trinsic.Services.VerifiableCredentials.V1.CreateProofRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[CreateProofRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest 'Trinsic.Services.VerifiableCredentials.V1.CreateProofRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_DocumentJsonFieldNumber'></a>
## CreateProofRequest.DocumentJsonFieldNumber Field
Field number for the "document_json" field.
```csharp
public const int DocumentJsonFieldNumber = 3;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
<a name='Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_ItemIdFieldNumber'></a>
## CreateProofRequest.ItemIdFieldNumber Field
Field number for the "item_id" field.
```csharp
public const int ItemIdFieldNumber = 2;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
<a name='Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_RevealDocumentJsonFieldNumber'></a>
## CreateProofRequest.RevealDocumentJsonFieldNumber Field
Field number for the "reveal_document_json" field.
```csharp
public const int RevealDocumentJsonFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
### Properties
<a name='Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_DocumentJson'></a>
## CreateProofRequest.DocumentJson Property
A document that contains a valid verifiable credential with an  
unbound signature. The proof will be derived from this document  
directly. The document will not be stored in the wallet.  
```csharp
public string DocumentJson { get; set; }
```
#### Property Value
[System.String](https://docs.microsoft.com/en-us/dotnet/api/System.String 'System.String')
  
<a name='Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_ItemId'></a>
## CreateProofRequest.ItemId Property
The item identifier that contains a record with a verifiable  
credential to be used for generating the proof.  
```csharp
public string ItemId { get; set; }
```
#### Property Value
[System.String](https://docs.microsoft.com/en-us/dotnet/api/System.String 'System.String')
  
<a name='Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_RevealDocumentJson'></a>
## CreateProofRequest.RevealDocumentJson Property
Optional document that describes which fields should be  
revealed in the generated proof. If specified, this document must be  
a valid JSON-LD frame.  
If this field is not specified, a default reveal document will be  
used and all fields in the signed document will be revealed  
```csharp
public string RevealDocumentJson { get; set; }
```
#### Property Value
[System.String](https://docs.microsoft.com/en-us/dotnet/api/System.String 'System.String')
  
  
<a name='Trinsic_Services_VerifiableCredentials_V1_CreateProofResponse'></a>
## CreateProofResponse Class
```csharp
public sealed class CreateProofResponse :
Google.Protobuf.IMessage<Trinsic.Services.VerifiableCredentials.V1.CreateProofResponse>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.VerifiableCredentials.V1.CreateProofResponse>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.VerifiableCredentials.V1.CreateProofResponse>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; CreateProofResponse  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[CreateProofResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CreateProofResponse 'Trinsic.Services.VerifiableCredentials.V1.CreateProofResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[CreateProofResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CreateProofResponse 'Trinsic.Services.VerifiableCredentials.V1.CreateProofResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[CreateProofResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CreateProofResponse 'Trinsic.Services.VerifiableCredentials.V1.CreateProofResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_VerifiableCredentials_V1_CreateProofResponse_ProofDocumentJsonFieldNumber'></a>
## CreateProofResponse.ProofDocumentJsonFieldNumber Field
Field number for the "proof_document_json" field.
```csharp
public const int ProofDocumentJsonFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
  
<a name='Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest'></a>
## IssueFromTemplateRequest Class
```csharp
public sealed class IssueFromTemplateRequest :
Google.Protobuf.IMessage<Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateRequest>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateRequest>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateRequest>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; IssueFromTemplateRequest  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[IssueFromTemplateRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest 'Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[IssueFromTemplateRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest 'Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[IssueFromTemplateRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest 'Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest_TemplateIdFieldNumber'></a>
## IssueFromTemplateRequest.TemplateIdFieldNumber Field
Field number for the "template_id" field.
```csharp
public const int TemplateIdFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
<a name='Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest_ValuesJsonFieldNumber'></a>
## IssueFromTemplateRequest.ValuesJsonFieldNumber Field
Field number for the "values_json" field.
```csharp
public const int ValuesJsonFieldNumber = 2;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
  
<a name='Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateResponse'></a>
## IssueFromTemplateResponse Class
```csharp
public sealed class IssueFromTemplateResponse :
Google.Protobuf.IMessage<Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateResponse>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateResponse>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateResponse>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; IssueFromTemplateResponse  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[IssueFromTemplateResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateResponse 'Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[IssueFromTemplateResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateResponse 'Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[IssueFromTemplateResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateResponse 'Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateResponse_DocumentJsonFieldNumber'></a>
## IssueFromTemplateResponse.DocumentJsonFieldNumber Field
Field number for the "document_json" field.
```csharp
public const int DocumentJsonFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
  
<a name='Trinsic_Services_VerifiableCredentials_V1_IssueRequest'></a>
## IssueRequest Class
```csharp
public sealed class IssueRequest :
Google.Protobuf.IMessage<Trinsic.Services.VerifiableCredentials.V1.IssueRequest>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.VerifiableCredentials.V1.IssueRequest>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.VerifiableCredentials.V1.IssueRequest>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; IssueRequest  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[IssueRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueRequest 'Trinsic.Services.VerifiableCredentials.V1.IssueRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[IssueRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueRequest 'Trinsic.Services.VerifiableCredentials.V1.IssueRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[IssueRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueRequest 'Trinsic.Services.VerifiableCredentials.V1.IssueRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_VerifiableCredentials_V1_IssueRequest_DocumentJsonFieldNumber'></a>
## IssueRequest.DocumentJsonFieldNumber Field
Field number for the "document_json" field.
```csharp
public const int DocumentJsonFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
  
<a name='Trinsic_Services_VerifiableCredentials_V1_IssueResponse'></a>
## IssueResponse Class
```csharp
public sealed class IssueResponse :
Google.Protobuf.IMessage<Trinsic.Services.VerifiableCredentials.V1.IssueResponse>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.VerifiableCredentials.V1.IssueResponse>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.VerifiableCredentials.V1.IssueResponse>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; IssueResponse  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[IssueResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueResponse 'Trinsic.Services.VerifiableCredentials.V1.IssueResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[IssueResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueResponse 'Trinsic.Services.VerifiableCredentials.V1.IssueResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[IssueResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueResponse 'Trinsic.Services.VerifiableCredentials.V1.IssueResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_VerifiableCredentials_V1_IssueResponse_SignedDocumentJsonFieldNumber'></a>
## IssueResponse.SignedDocumentJsonFieldNumber Field
Field number for the "signed_document_json" field.
```csharp
public const int SignedDocumentJsonFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
  
<a name='Trinsic_Services_VerifiableCredentials_V1_SendRequest'></a>
## SendRequest Class
```csharp
public sealed class SendRequest :
Google.Protobuf.IMessage<Trinsic.Services.VerifiableCredentials.V1.SendRequest>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.VerifiableCredentials.V1.SendRequest>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.VerifiableCredentials.V1.SendRequest>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; SendRequest  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[SendRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_SendRequest 'Trinsic.Services.VerifiableCredentials.V1.SendRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[SendRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_SendRequest 'Trinsic.Services.VerifiableCredentials.V1.SendRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[SendRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_SendRequest 'Trinsic.Services.VerifiableCredentials.V1.SendRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_VerifiableCredentials_V1_SendRequest_DidcommInvitationJsonFieldNumber'></a>
## SendRequest.DidcommInvitationJsonFieldNumber Field
Field number for the "didcomm_invitation_json" field.
```csharp
public const int DidcommInvitationJsonFieldNumber = 3;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
<a name='Trinsic_Services_VerifiableCredentials_V1_SendRequest_DidUriFieldNumber'></a>
## SendRequest.DidUriFieldNumber Field
Field number for the "did_uri" field.
```csharp
public const int DidUriFieldNumber = 2;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
<a name='Trinsic_Services_VerifiableCredentials_V1_SendRequest_DocumentJsonFieldNumber'></a>
## SendRequest.DocumentJsonFieldNumber Field
Field number for the "document_json" field.
```csharp
public const int DocumentJsonFieldNumber = 100;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
<a name='Trinsic_Services_VerifiableCredentials_V1_SendRequest_EmailFieldNumber'></a>
## SendRequest.EmailFieldNumber Field
Field number for the "email" field.
```csharp
public const int EmailFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
  
<a name='Trinsic_Services_VerifiableCredentials_V1_SendResponse'></a>
## SendResponse Class
```csharp
public sealed class SendResponse :
Google.Protobuf.IMessage<Trinsic.Services.VerifiableCredentials.V1.SendResponse>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.VerifiableCredentials.V1.SendResponse>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.VerifiableCredentials.V1.SendResponse>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; SendResponse  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[SendResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_SendResponse 'Trinsic.Services.VerifiableCredentials.V1.SendResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[SendResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_SendResponse 'Trinsic.Services.VerifiableCredentials.V1.SendResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[SendResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_SendResponse 'Trinsic.Services.VerifiableCredentials.V1.SendResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_VerifiableCredentials_V1_SendResponse_StatusFieldNumber'></a>
## SendResponse.StatusFieldNumber Field
Field number for the "status" field.
```csharp
public const int StatusFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
  
<a name='Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest'></a>
## UpdateStatusRequest Class
request object to update the status of the revocation entry  
```csharp
public sealed class UpdateStatusRequest :
Google.Protobuf.IMessage<Trinsic.Services.VerifiableCredentials.V1.UpdateStatusRequest>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.VerifiableCredentials.V1.UpdateStatusRequest>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.VerifiableCredentials.V1.UpdateStatusRequest>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; UpdateStatusRequest  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[UpdateStatusRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest 'Trinsic.Services.VerifiableCredentials.V1.UpdateStatusRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[UpdateStatusRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest 'Trinsic.Services.VerifiableCredentials.V1.UpdateStatusRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[UpdateStatusRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest 'Trinsic.Services.VerifiableCredentials.V1.UpdateStatusRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_CredentialStatusIdFieldNumber'></a>
## UpdateStatusRequest.CredentialStatusIdFieldNumber Field
Field number for the "credential_status_id" field.
```csharp
public const int CredentialStatusIdFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
<a name='Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_RevokedFieldNumber'></a>
## UpdateStatusRequest.RevokedFieldNumber Field
Field number for the "revoked" field.
```csharp
public const int RevokedFieldNumber = 2;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
### Properties
<a name='Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_CredentialStatusId'></a>
## UpdateStatusRequest.CredentialStatusId Property
the credential status id  
```csharp
public string CredentialStatusId { get; set; }
```
#### Property Value
[System.String](https://docs.microsoft.com/en-us/dotnet/api/System.String 'System.String')
  
<a name='Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_Revoked'></a>
## UpdateStatusRequest.Revoked Property
indicates if the status is revoked  
```csharp
public bool Revoked { get; set; }
```
#### Property Value
[System.Boolean](https://docs.microsoft.com/en-us/dotnet/api/System.Boolean 'System.Boolean')
  
  
<a name='Trinsic_Services_VerifiableCredentials_V1_UpdateStatusResponse'></a>
## UpdateStatusResponse Class
response object for update of status of revocation entry  
```csharp
public sealed class UpdateStatusResponse :
Google.Protobuf.IMessage<Trinsic.Services.VerifiableCredentials.V1.UpdateStatusResponse>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.VerifiableCredentials.V1.UpdateStatusResponse>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.VerifiableCredentials.V1.UpdateStatusResponse>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; UpdateStatusResponse  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[UpdateStatusResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_UpdateStatusResponse 'Trinsic.Services.VerifiableCredentials.V1.UpdateStatusResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[UpdateStatusResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_UpdateStatusResponse 'Trinsic.Services.VerifiableCredentials.V1.UpdateStatusResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[UpdateStatusResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_UpdateStatusResponse 'Trinsic.Services.VerifiableCredentials.V1.UpdateStatusResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_VerifiableCredentials_V1_UpdateStatusResponse_StatusFieldNumber'></a>
## UpdateStatusResponse.StatusFieldNumber Field
Field number for the "status" field.
```csharp
public const int StatusFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential'></a>
## VerifiableCredential Class
```csharp
public static class VerifiableCredential
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; VerifiableCredential  
### Properties
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_Descriptor'></a>
## VerifiableCredential.Descriptor Property
Service descriptor
```csharp
public static Google.Protobuf.Reflection.ServiceDescriptor Descriptor { get; }
```
#### Property Value
[Google.Protobuf.Reflection.ServiceDescriptor](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.Reflection.ServiceDescriptor 'Google.Protobuf.Reflection.ServiceDescriptor')
  
### Methods
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_BindService(Grpc_Core_ServiceBinderBase_Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase)'></a>
## VerifiableCredential.BindService(ServiceBinderBase, VerifiableCredentialBase) Method
Register service method with a service binder with or without implementation. Useful when customizing the  service binding logic.  
            Note: this method is part of an experimental API that can change or be removed without any prior notice.
```csharp
public static void BindService(Grpc.Core.ServiceBinderBase serviceBinder, Trinsic.Services.VerifiableCredentials.V1.VerifiableCredential.VerifiableCredentialBase serviceImpl);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_BindService(Grpc_Core_ServiceBinderBase_Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase)_serviceBinder'></a>
`serviceBinder` [Grpc.Core.ServiceBinderBase](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ServiceBinderBase 'Grpc.Core.ServiceBinderBase')  
Service methods will be bound by calling `AddMethod` on this object.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_BindService(Grpc_Core_ServiceBinderBase_Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase)_serviceImpl'></a>
`serviceImpl` [VerifiableCredentialBase](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase 'Trinsic.Services.VerifiableCredentials.V1.VerifiableCredential.VerifiableCredentialBase')  
An object implementing the server-side handling logic.
  
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_BindService(Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase)'></a>
## VerifiableCredential.BindService(VerifiableCredentialBase) Method
Creates service definition that can be registered with a server
```csharp
public static Grpc.Core.ServerServiceDefinition BindService(Trinsic.Services.VerifiableCredentials.V1.VerifiableCredential.VerifiableCredentialBase serviceImpl);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_BindService(Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase)_serviceImpl'></a>
`serviceImpl` [VerifiableCredentialBase](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase 'Trinsic.Services.VerifiableCredentials.V1.VerifiableCredential.VerifiableCredentialBase')  
An object implementing the server-side handling logic.
  
#### Returns
[Grpc.Core.ServerServiceDefinition](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ServerServiceDefinition 'Grpc.Core.ServerServiceDefinition')  
  
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase'></a>
## VerifiableCredential.VerifiableCredentialBase Class
Base class for server-side implementations of VerifiableCredential
```csharp
public abstract class VerifiableCredential.VerifiableCredentialBase
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; VerifiableCredentialBase  
### Methods
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase_CheckStatus(Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest_Grpc_Core_ServerCallContext)'></a>
## VerifiableCredential.VerifiableCredentialBase.CheckStatus(CheckStatusRequest, ServerCallContext) Method
Check credential status by setting the revocation value  
```csharp
public virtual System.Threading.Tasks.Task<Trinsic.Services.VerifiableCredentials.V1.CheckStatusResponse> CheckStatus(Trinsic.Services.VerifiableCredentials.V1.CheckStatusRequest request, Grpc.Core.ServerCallContext context);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase_CheckStatus(Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest_Grpc_Core_ServerCallContext)_request'></a>
`request` [CheckStatusRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest 'Trinsic.Services.VerifiableCredentials.V1.CheckStatusRequest')  
The request received from the client.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase_CheckStatus(Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest_Grpc_Core_ServerCallContext)_context'></a>
`context` [Grpc.Core.ServerCallContext](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ServerCallContext 'Grpc.Core.ServerCallContext')  
The context of the server-side call handler being invoked.
  
#### Returns
[System.Threading.Tasks.Task&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')[CheckStatusResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CheckStatusResponse 'Trinsic.Services.VerifiableCredentials.V1.CheckStatusResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')  
The response to send back to the client (wrapped by a task).
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase_CreateProof(Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_Grpc_Core_ServerCallContext)'></a>
## VerifiableCredential.VerifiableCredentialBase.CreateProof(CreateProofRequest, ServerCallContext) Method
Create a proof from a signed document that is a valid  
verifiable credential and contains a signature from which a proof can be derived.  
```csharp
public virtual System.Threading.Tasks.Task<Trinsic.Services.VerifiableCredentials.V1.CreateProofResponse> CreateProof(Trinsic.Services.VerifiableCredentials.V1.CreateProofRequest request, Grpc.Core.ServerCallContext context);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase_CreateProof(Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_Grpc_Core_ServerCallContext)_request'></a>
`request` [CreateProofRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest 'Trinsic.Services.VerifiableCredentials.V1.CreateProofRequest')  
The request received from the client.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase_CreateProof(Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_Grpc_Core_ServerCallContext)_context'></a>
`context` [Grpc.Core.ServerCallContext](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ServerCallContext 'Grpc.Core.ServerCallContext')  
The context of the server-side call handler being invoked.
  
#### Returns
[System.Threading.Tasks.Task&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')[CreateProofResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CreateProofResponse 'Trinsic.Services.VerifiableCredentials.V1.CreateProofResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')  
The response to send back to the client (wrapped by a task).
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase_Issue(Trinsic_Services_VerifiableCredentials_V1_IssueRequest_Grpc_Core_ServerCallContext)'></a>
## VerifiableCredential.VerifiableCredentialBase.Issue(IssueRequest, ServerCallContext) Method
Sign and issue a verifiable credential from a submitted document.  
The document must be a valid JSON-LD document.  
```csharp
public virtual System.Threading.Tasks.Task<Trinsic.Services.VerifiableCredentials.V1.IssueResponse> Issue(Trinsic.Services.VerifiableCredentials.V1.IssueRequest request, Grpc.Core.ServerCallContext context);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase_Issue(Trinsic_Services_VerifiableCredentials_V1_IssueRequest_Grpc_Core_ServerCallContext)_request'></a>
`request` [IssueRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueRequest 'Trinsic.Services.VerifiableCredentials.V1.IssueRequest')  
The request received from the client.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase_Issue(Trinsic_Services_VerifiableCredentials_V1_IssueRequest_Grpc_Core_ServerCallContext)_context'></a>
`context` [Grpc.Core.ServerCallContext](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ServerCallContext 'Grpc.Core.ServerCallContext')  
The context of the server-side call handler being invoked.
  
#### Returns
[System.Threading.Tasks.Task&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')[IssueResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueResponse 'Trinsic.Services.VerifiableCredentials.V1.IssueResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')  
The response to send back to the client (wrapped by a task).
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase_IssueFromTemplate(Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest_Grpc_Core_ServerCallContext)'></a>
## VerifiableCredential.VerifiableCredentialBase.IssueFromTemplate(IssueFromTemplateRequest, ServerCallContext) Method
Sign and issue a verifiable credential from a pre-defined template.  
This process will also add schema validation and   
revocation registry entry in the credential.  
```csharp
public virtual System.Threading.Tasks.Task<Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateResponse> IssueFromTemplate(Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateRequest request, Grpc.Core.ServerCallContext context);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase_IssueFromTemplate(Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest_Grpc_Core_ServerCallContext)_request'></a>
`request` [IssueFromTemplateRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest 'Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateRequest')  
The request received from the client.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase_IssueFromTemplate(Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest_Grpc_Core_ServerCallContext)_context'></a>
`context` [Grpc.Core.ServerCallContext](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ServerCallContext 'Grpc.Core.ServerCallContext')  
The context of the server-side call handler being invoked.
  
#### Returns
[System.Threading.Tasks.Task&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')[IssueFromTemplateResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateResponse 'Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')  
The response to send back to the client (wrapped by a task).
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase_Send(Trinsic_Services_VerifiableCredentials_V1_SendRequest_Grpc_Core_ServerCallContext)'></a>
## VerifiableCredential.VerifiableCredentialBase.Send(SendRequest, ServerCallContext) Method
Sends a document directly to a user's email within the given ecosystem  
```csharp
public virtual System.Threading.Tasks.Task<Trinsic.Services.VerifiableCredentials.V1.SendResponse> Send(Trinsic.Services.VerifiableCredentials.V1.SendRequest request, Grpc.Core.ServerCallContext context);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase_Send(Trinsic_Services_VerifiableCredentials_V1_SendRequest_Grpc_Core_ServerCallContext)_request'></a>
`request` [SendRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_SendRequest 'Trinsic.Services.VerifiableCredentials.V1.SendRequest')  
The request received from the client.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase_Send(Trinsic_Services_VerifiableCredentials_V1_SendRequest_Grpc_Core_ServerCallContext)_context'></a>
`context` [Grpc.Core.ServerCallContext](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ServerCallContext 'Grpc.Core.ServerCallContext')  
The context of the server-side call handler being invoked.
  
#### Returns
[System.Threading.Tasks.Task&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')[SendResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_SendResponse 'Trinsic.Services.VerifiableCredentials.V1.SendResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')  
The response to send back to the client (wrapped by a task).
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase_UpdateStatus(Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_Grpc_Core_ServerCallContext)'></a>
## VerifiableCredential.VerifiableCredentialBase.UpdateStatus(UpdateStatusRequest, ServerCallContext) Method
Update credential status by setting the revocation value  
```csharp
public virtual System.Threading.Tasks.Task<Trinsic.Services.VerifiableCredentials.V1.UpdateStatusResponse> UpdateStatus(Trinsic.Services.VerifiableCredentials.V1.UpdateStatusRequest request, Grpc.Core.ServerCallContext context);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase_UpdateStatus(Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_Grpc_Core_ServerCallContext)_request'></a>
`request` [UpdateStatusRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest 'Trinsic.Services.VerifiableCredentials.V1.UpdateStatusRequest')  
The request received from the client.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase_UpdateStatus(Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_Grpc_Core_ServerCallContext)_context'></a>
`context` [Grpc.Core.ServerCallContext](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ServerCallContext 'Grpc.Core.ServerCallContext')  
The context of the server-side call handler being invoked.
  
#### Returns
[System.Threading.Tasks.Task&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')[UpdateStatusResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_UpdateStatusResponse 'Trinsic.Services.VerifiableCredentials.V1.UpdateStatusResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')  
The response to send back to the client (wrapped by a task).
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase_VerifyProof(Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest_Grpc_Core_ServerCallContext)'></a>
## VerifiableCredential.VerifiableCredentialBase.VerifyProof(VerifyProofRequest, ServerCallContext) Method
Verifies a proof by checking the signature value, and if possible schema validation,  
revocation status, and issuer status against a trust registry  
```csharp
public virtual System.Threading.Tasks.Task<Trinsic.Services.VerifiableCredentials.V1.VerifyProofResponse> VerifyProof(Trinsic.Services.VerifiableCredentials.V1.VerifyProofRequest request, Grpc.Core.ServerCallContext context);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase_VerifyProof(Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest_Grpc_Core_ServerCallContext)_request'></a>
`request` [VerifyProofRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest 'Trinsic.Services.VerifiableCredentials.V1.VerifyProofRequest')  
The request received from the client.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialBase_VerifyProof(Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest_Grpc_Core_ServerCallContext)_context'></a>
`context` [Grpc.Core.ServerCallContext](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ServerCallContext 'Grpc.Core.ServerCallContext')  
The context of the server-side call handler being invoked.
  
#### Returns
[System.Threading.Tasks.Task&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')[VerifyProofResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_VerifyProofResponse 'Trinsic.Services.VerifiableCredentials.V1.VerifyProofResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')  
The response to send back to the client (wrapped by a task).
  
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient'></a>
## VerifiableCredential.VerifiableCredentialClient Class
Client for VerifiableCredential
```csharp
public class VerifiableCredential.VerifiableCredentialClient : Grpc.Core.ClientBase<Trinsic.Services.VerifiableCredentials.V1.VerifiableCredential.VerifiableCredentialClient>
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; [Grpc.Core.ClientBase](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ClientBase 'Grpc.Core.ClientBase') &#129106; [Grpc.Core.ClientBase&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ClientBase-1 'Grpc.Core.ClientBase`1')[VerifiableCredentialClient](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient 'Trinsic.Services.VerifiableCredentials.V1.VerifiableCredential.VerifiableCredentialClient')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ClientBase-1 'Grpc.Core.ClientBase`1') &#129106; VerifiableCredentialClient  
### Constructors
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifiableCredentialClient()'></a>
## VerifiableCredential.VerifiableCredentialClient.VerifiableCredentialClient() Constructor
Protected parameterless constructor to allow creation of test doubles.
```csharp
protected VerifiableCredentialClient();
```
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifiableCredentialClient(Grpc_Core_CallInvoker)'></a>
## VerifiableCredential.VerifiableCredentialClient.VerifiableCredentialClient(CallInvoker) Constructor
Creates a new client for VerifiableCredential that uses a custom `CallInvoker`.
```csharp
public VerifiableCredentialClient(Grpc.Core.CallInvoker callInvoker);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifiableCredentialClient(Grpc_Core_CallInvoker)_callInvoker'></a>
`callInvoker` [Grpc.Core.CallInvoker](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallInvoker 'Grpc.Core.CallInvoker')  
The callInvoker to use to make remote calls.
  
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifiableCredentialClient(Grpc_Core_ChannelBase)'></a>
## VerifiableCredential.VerifiableCredentialClient.VerifiableCredentialClient(ChannelBase) Constructor
Creates a new client for VerifiableCredential
```csharp
public VerifiableCredentialClient(Grpc.Core.ChannelBase channel);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifiableCredentialClient(Grpc_Core_ChannelBase)_channel'></a>
`channel` [Grpc.Core.ChannelBase](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ChannelBase 'Grpc.Core.ChannelBase')  
The channel to use to make remote calls.
  
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifiableCredentialClient(Grpc_Core_ClientBase_ClientBaseConfiguration)'></a>
## VerifiableCredential.VerifiableCredentialClient.VerifiableCredentialClient(ClientBaseConfiguration) Constructor
Protected constructor to allow creation of configured clients.
```csharp
protected VerifiableCredentialClient(Grpc.Core.ClientBase.ClientBaseConfiguration configuration);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifiableCredentialClient(Grpc_Core_ClientBase_ClientBaseConfiguration)_configuration'></a>
`configuration` [Grpc.Core.ClientBase.ClientBaseConfiguration](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ClientBase.ClientBaseConfiguration 'Grpc.Core.ClientBase.ClientBaseConfiguration')  
The client configuration.
  
  
### Methods
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CheckStatus(Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest_Grpc_Core_CallOptions)'></a>
## VerifiableCredential.VerifiableCredentialClient.CheckStatus(CheckStatusRequest, CallOptions) Method
Check credential status by setting the revocation value  
```csharp
public virtual Trinsic.Services.VerifiableCredentials.V1.CheckStatusResponse CheckStatus(Trinsic.Services.VerifiableCredentials.V1.CheckStatusRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CheckStatus(Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest_Grpc_Core_CallOptions)_request'></a>
`request` [CheckStatusRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest 'Trinsic.Services.VerifiableCredentials.V1.CheckStatusRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CheckStatus(Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[CheckStatusResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CheckStatusResponse 'Trinsic.Services.VerifiableCredentials.V1.CheckStatusResponse')  
The response received from the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CheckStatus(Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## VerifiableCredential.VerifiableCredentialClient.CheckStatus(CheckStatusRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Check credential status by setting the revocation value  
```csharp
public virtual Trinsic.Services.VerifiableCredentials.V1.CheckStatusResponse CheckStatus(Trinsic.Services.VerifiableCredentials.V1.CheckStatusRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CheckStatus(Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [CheckStatusRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest 'Trinsic.Services.VerifiableCredentials.V1.CheckStatusRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CheckStatus(Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CheckStatus(Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CheckStatus(Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[CheckStatusResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CheckStatusResponse 'Trinsic.Services.VerifiableCredentials.V1.CheckStatusResponse')  
The response received from the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CheckStatusAsync(Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest_Grpc_Core_CallOptions)'></a>
## VerifiableCredential.VerifiableCredentialClient.CheckStatusAsync(CheckStatusRequest, CallOptions) Method
Check credential status by setting the revocation value  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.VerifiableCredentials.V1.CheckStatusResponse> CheckStatusAsync(Trinsic.Services.VerifiableCredentials.V1.CheckStatusRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CheckStatusAsync(Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest_Grpc_Core_CallOptions)_request'></a>
`request` [CheckStatusRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest 'Trinsic.Services.VerifiableCredentials.V1.CheckStatusRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CheckStatusAsync(Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[CheckStatusResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CheckStatusResponse 'Trinsic.Services.VerifiableCredentials.V1.CheckStatusResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CheckStatusAsync(Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## VerifiableCredential.VerifiableCredentialClient.CheckStatusAsync(CheckStatusRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Check credential status by setting the revocation value  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.VerifiableCredentials.V1.CheckStatusResponse> CheckStatusAsync(Trinsic.Services.VerifiableCredentials.V1.CheckStatusRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CheckStatusAsync(Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [CheckStatusRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest 'Trinsic.Services.VerifiableCredentials.V1.CheckStatusRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CheckStatusAsync(Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CheckStatusAsync(Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CheckStatusAsync(Trinsic_Services_VerifiableCredentials_V1_CheckStatusRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[CheckStatusResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CheckStatusResponse 'Trinsic.Services.VerifiableCredentials.V1.CheckStatusResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CreateProof(Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_Grpc_Core_CallOptions)'></a>
## VerifiableCredential.VerifiableCredentialClient.CreateProof(CreateProofRequest, CallOptions) Method
Create a proof from a signed document that is a valid  
verifiable credential and contains a signature from which a proof can be derived.  
```csharp
public virtual Trinsic.Services.VerifiableCredentials.V1.CreateProofResponse CreateProof(Trinsic.Services.VerifiableCredentials.V1.CreateProofRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CreateProof(Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_Grpc_Core_CallOptions)_request'></a>
`request` [CreateProofRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest 'Trinsic.Services.VerifiableCredentials.V1.CreateProofRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CreateProof(Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[CreateProofResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CreateProofResponse 'Trinsic.Services.VerifiableCredentials.V1.CreateProofResponse')  
The response received from the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CreateProof(Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## VerifiableCredential.VerifiableCredentialClient.CreateProof(CreateProofRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Create a proof from a signed document that is a valid  
verifiable credential and contains a signature from which a proof can be derived.  
```csharp
public virtual Trinsic.Services.VerifiableCredentials.V1.CreateProofResponse CreateProof(Trinsic.Services.VerifiableCredentials.V1.CreateProofRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CreateProof(Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [CreateProofRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest 'Trinsic.Services.VerifiableCredentials.V1.CreateProofRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CreateProof(Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CreateProof(Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CreateProof(Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[CreateProofResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CreateProofResponse 'Trinsic.Services.VerifiableCredentials.V1.CreateProofResponse')  
The response received from the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CreateProofAsync(Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_Grpc_Core_CallOptions)'></a>
## VerifiableCredential.VerifiableCredentialClient.CreateProofAsync(CreateProofRequest, CallOptions) Method
Create a proof from a signed document that is a valid  
verifiable credential and contains a signature from which a proof can be derived.  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.VerifiableCredentials.V1.CreateProofResponse> CreateProofAsync(Trinsic.Services.VerifiableCredentials.V1.CreateProofRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CreateProofAsync(Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_Grpc_Core_CallOptions)_request'></a>
`request` [CreateProofRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest 'Trinsic.Services.VerifiableCredentials.V1.CreateProofRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CreateProofAsync(Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[CreateProofResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CreateProofResponse 'Trinsic.Services.VerifiableCredentials.V1.CreateProofResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CreateProofAsync(Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## VerifiableCredential.VerifiableCredentialClient.CreateProofAsync(CreateProofRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Create a proof from a signed document that is a valid  
verifiable credential and contains a signature from which a proof can be derived.  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.VerifiableCredentials.V1.CreateProofResponse> CreateProofAsync(Trinsic.Services.VerifiableCredentials.V1.CreateProofRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CreateProofAsync(Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [CreateProofRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest 'Trinsic.Services.VerifiableCredentials.V1.CreateProofRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CreateProofAsync(Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CreateProofAsync(Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_CreateProofAsync(Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[CreateProofResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_CreateProofResponse 'Trinsic.Services.VerifiableCredentials.V1.CreateProofResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_Issue(Trinsic_Services_VerifiableCredentials_V1_IssueRequest_Grpc_Core_CallOptions)'></a>
## VerifiableCredential.VerifiableCredentialClient.Issue(IssueRequest, CallOptions) Method
Sign and issue a verifiable credential from a submitted document.  
The document must be a valid JSON-LD document.  
```csharp
public virtual Trinsic.Services.VerifiableCredentials.V1.IssueResponse Issue(Trinsic.Services.VerifiableCredentials.V1.IssueRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_Issue(Trinsic_Services_VerifiableCredentials_V1_IssueRequest_Grpc_Core_CallOptions)_request'></a>
`request` [IssueRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueRequest 'Trinsic.Services.VerifiableCredentials.V1.IssueRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_Issue(Trinsic_Services_VerifiableCredentials_V1_IssueRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[IssueResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueResponse 'Trinsic.Services.VerifiableCredentials.V1.IssueResponse')  
The response received from the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_Issue(Trinsic_Services_VerifiableCredentials_V1_IssueRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## VerifiableCredential.VerifiableCredentialClient.Issue(IssueRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Sign and issue a verifiable credential from a submitted document.  
The document must be a valid JSON-LD document.  
```csharp
public virtual Trinsic.Services.VerifiableCredentials.V1.IssueResponse Issue(Trinsic.Services.VerifiableCredentials.V1.IssueRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_Issue(Trinsic_Services_VerifiableCredentials_V1_IssueRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [IssueRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueRequest 'Trinsic.Services.VerifiableCredentials.V1.IssueRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_Issue(Trinsic_Services_VerifiableCredentials_V1_IssueRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_Issue(Trinsic_Services_VerifiableCredentials_V1_IssueRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_Issue(Trinsic_Services_VerifiableCredentials_V1_IssueRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[IssueResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueResponse 'Trinsic.Services.VerifiableCredentials.V1.IssueResponse')  
The response received from the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueAsync(Trinsic_Services_VerifiableCredentials_V1_IssueRequest_Grpc_Core_CallOptions)'></a>
## VerifiableCredential.VerifiableCredentialClient.IssueAsync(IssueRequest, CallOptions) Method
Sign and issue a verifiable credential from a submitted document.  
The document must be a valid JSON-LD document.  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.VerifiableCredentials.V1.IssueResponse> IssueAsync(Trinsic.Services.VerifiableCredentials.V1.IssueRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueAsync(Trinsic_Services_VerifiableCredentials_V1_IssueRequest_Grpc_Core_CallOptions)_request'></a>
`request` [IssueRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueRequest 'Trinsic.Services.VerifiableCredentials.V1.IssueRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueAsync(Trinsic_Services_VerifiableCredentials_V1_IssueRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[IssueResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueResponse 'Trinsic.Services.VerifiableCredentials.V1.IssueResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueAsync(Trinsic_Services_VerifiableCredentials_V1_IssueRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## VerifiableCredential.VerifiableCredentialClient.IssueAsync(IssueRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Sign and issue a verifiable credential from a submitted document.  
The document must be a valid JSON-LD document.  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.VerifiableCredentials.V1.IssueResponse> IssueAsync(Trinsic.Services.VerifiableCredentials.V1.IssueRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueAsync(Trinsic_Services_VerifiableCredentials_V1_IssueRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [IssueRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueRequest 'Trinsic.Services.VerifiableCredentials.V1.IssueRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueAsync(Trinsic_Services_VerifiableCredentials_V1_IssueRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueAsync(Trinsic_Services_VerifiableCredentials_V1_IssueRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueAsync(Trinsic_Services_VerifiableCredentials_V1_IssueRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[IssueResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueResponse 'Trinsic.Services.VerifiableCredentials.V1.IssueResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueFromTemplate(Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest_Grpc_Core_CallOptions)'></a>
## VerifiableCredential.VerifiableCredentialClient.IssueFromTemplate(IssueFromTemplateRequest, CallOptions) Method
Sign and issue a verifiable credential from a pre-defined template.  
This process will also add schema validation and   
revocation registry entry in the credential.  
```csharp
public virtual Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateResponse IssueFromTemplate(Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueFromTemplate(Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest_Grpc_Core_CallOptions)_request'></a>
`request` [IssueFromTemplateRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest 'Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueFromTemplate(Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[IssueFromTemplateResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateResponse 'Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateResponse')  
The response received from the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueFromTemplate(Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## VerifiableCredential.VerifiableCredentialClient.IssueFromTemplate(IssueFromTemplateRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Sign and issue a verifiable credential from a pre-defined template.  
This process will also add schema validation and   
revocation registry entry in the credential.  
```csharp
public virtual Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateResponse IssueFromTemplate(Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueFromTemplate(Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [IssueFromTemplateRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest 'Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueFromTemplate(Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueFromTemplate(Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueFromTemplate(Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[IssueFromTemplateResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateResponse 'Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateResponse')  
The response received from the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueFromTemplateAsync(Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest_Grpc_Core_CallOptions)'></a>
## VerifiableCredential.VerifiableCredentialClient.IssueFromTemplateAsync(IssueFromTemplateRequest, CallOptions) Method
Sign and issue a verifiable credential from a pre-defined template.  
This process will also add schema validation and   
revocation registry entry in the credential.  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateResponse> IssueFromTemplateAsync(Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueFromTemplateAsync(Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest_Grpc_Core_CallOptions)_request'></a>
`request` [IssueFromTemplateRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest 'Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueFromTemplateAsync(Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[IssueFromTemplateResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateResponse 'Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueFromTemplateAsync(Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## VerifiableCredential.VerifiableCredentialClient.IssueFromTemplateAsync(IssueFromTemplateRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Sign and issue a verifiable credential from a pre-defined template.  
This process will also add schema validation and   
revocation registry entry in the credential.  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateResponse> IssueFromTemplateAsync(Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueFromTemplateAsync(Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [IssueFromTemplateRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest 'Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueFromTemplateAsync(Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueFromTemplateAsync(Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_IssueFromTemplateAsync(Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[IssueFromTemplateResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_IssueFromTemplateResponse 'Trinsic.Services.VerifiableCredentials.V1.IssueFromTemplateResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_NewInstance(Grpc_Core_ClientBase_ClientBaseConfiguration)'></a>
## VerifiableCredential.VerifiableCredentialClient.NewInstance(ClientBaseConfiguration) Method
Creates a new instance of client from given `ClientBaseConfiguration`.
```csharp
protected override Trinsic.Services.VerifiableCredentials.V1.VerifiableCredential.VerifiableCredentialClient NewInstance(Grpc.Core.ClientBase.ClientBaseConfiguration configuration);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_NewInstance(Grpc_Core_ClientBase_ClientBaseConfiguration)_configuration'></a>
`configuration` [Grpc.Core.ClientBase.ClientBaseConfiguration](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ClientBase.ClientBaseConfiguration 'Grpc.Core.ClientBase.ClientBaseConfiguration')  
  
#### Returns
[VerifiableCredentialClient](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient 'Trinsic.Services.VerifiableCredentials.V1.VerifiableCredential.VerifiableCredentialClient')  
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_Send(Trinsic_Services_VerifiableCredentials_V1_SendRequest_Grpc_Core_CallOptions)'></a>
## VerifiableCredential.VerifiableCredentialClient.Send(SendRequest, CallOptions) Method
Sends a document directly to a user's email within the given ecosystem  
```csharp
public virtual Trinsic.Services.VerifiableCredentials.V1.SendResponse Send(Trinsic.Services.VerifiableCredentials.V1.SendRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_Send(Trinsic_Services_VerifiableCredentials_V1_SendRequest_Grpc_Core_CallOptions)_request'></a>
`request` [SendRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_SendRequest 'Trinsic.Services.VerifiableCredentials.V1.SendRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_Send(Trinsic_Services_VerifiableCredentials_V1_SendRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[SendResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_SendResponse 'Trinsic.Services.VerifiableCredentials.V1.SendResponse')  
The response received from the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_Send(Trinsic_Services_VerifiableCredentials_V1_SendRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## VerifiableCredential.VerifiableCredentialClient.Send(SendRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Sends a document directly to a user's email within the given ecosystem  
```csharp
public virtual Trinsic.Services.VerifiableCredentials.V1.SendResponse Send(Trinsic.Services.VerifiableCredentials.V1.SendRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_Send(Trinsic_Services_VerifiableCredentials_V1_SendRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [SendRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_SendRequest 'Trinsic.Services.VerifiableCredentials.V1.SendRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_Send(Trinsic_Services_VerifiableCredentials_V1_SendRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_Send(Trinsic_Services_VerifiableCredentials_V1_SendRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_Send(Trinsic_Services_VerifiableCredentials_V1_SendRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[SendResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_SendResponse 'Trinsic.Services.VerifiableCredentials.V1.SendResponse')  
The response received from the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_SendAsync(Trinsic_Services_VerifiableCredentials_V1_SendRequest_Grpc_Core_CallOptions)'></a>
## VerifiableCredential.VerifiableCredentialClient.SendAsync(SendRequest, CallOptions) Method
Sends a document directly to a user's email within the given ecosystem  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.VerifiableCredentials.V1.SendResponse> SendAsync(Trinsic.Services.VerifiableCredentials.V1.SendRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_SendAsync(Trinsic_Services_VerifiableCredentials_V1_SendRequest_Grpc_Core_CallOptions)_request'></a>
`request` [SendRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_SendRequest 'Trinsic.Services.VerifiableCredentials.V1.SendRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_SendAsync(Trinsic_Services_VerifiableCredentials_V1_SendRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[SendResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_SendResponse 'Trinsic.Services.VerifiableCredentials.V1.SendResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_SendAsync(Trinsic_Services_VerifiableCredentials_V1_SendRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## VerifiableCredential.VerifiableCredentialClient.SendAsync(SendRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Sends a document directly to a user's email within the given ecosystem  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.VerifiableCredentials.V1.SendResponse> SendAsync(Trinsic.Services.VerifiableCredentials.V1.SendRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_SendAsync(Trinsic_Services_VerifiableCredentials_V1_SendRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [SendRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_SendRequest 'Trinsic.Services.VerifiableCredentials.V1.SendRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_SendAsync(Trinsic_Services_VerifiableCredentials_V1_SendRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_SendAsync(Trinsic_Services_VerifiableCredentials_V1_SendRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_SendAsync(Trinsic_Services_VerifiableCredentials_V1_SendRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[SendResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_SendResponse 'Trinsic.Services.VerifiableCredentials.V1.SendResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_UpdateStatus(Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_Grpc_Core_CallOptions)'></a>
## VerifiableCredential.VerifiableCredentialClient.UpdateStatus(UpdateStatusRequest, CallOptions) Method
Update credential status by setting the revocation value  
```csharp
public virtual Trinsic.Services.VerifiableCredentials.V1.UpdateStatusResponse UpdateStatus(Trinsic.Services.VerifiableCredentials.V1.UpdateStatusRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_UpdateStatus(Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_Grpc_Core_CallOptions)_request'></a>
`request` [UpdateStatusRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest 'Trinsic.Services.VerifiableCredentials.V1.UpdateStatusRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_UpdateStatus(Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[UpdateStatusResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_UpdateStatusResponse 'Trinsic.Services.VerifiableCredentials.V1.UpdateStatusResponse')  
The response received from the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_UpdateStatus(Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## VerifiableCredential.VerifiableCredentialClient.UpdateStatus(UpdateStatusRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Update credential status by setting the revocation value  
```csharp
public virtual Trinsic.Services.VerifiableCredentials.V1.UpdateStatusResponse UpdateStatus(Trinsic.Services.VerifiableCredentials.V1.UpdateStatusRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_UpdateStatus(Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [UpdateStatusRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest 'Trinsic.Services.VerifiableCredentials.V1.UpdateStatusRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_UpdateStatus(Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_UpdateStatus(Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_UpdateStatus(Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[UpdateStatusResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_UpdateStatusResponse 'Trinsic.Services.VerifiableCredentials.V1.UpdateStatusResponse')  
The response received from the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_UpdateStatusAsync(Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_Grpc_Core_CallOptions)'></a>
## VerifiableCredential.VerifiableCredentialClient.UpdateStatusAsync(UpdateStatusRequest, CallOptions) Method
Update credential status by setting the revocation value  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.VerifiableCredentials.V1.UpdateStatusResponse> UpdateStatusAsync(Trinsic.Services.VerifiableCredentials.V1.UpdateStatusRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_UpdateStatusAsync(Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_Grpc_Core_CallOptions)_request'></a>
`request` [UpdateStatusRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest 'Trinsic.Services.VerifiableCredentials.V1.UpdateStatusRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_UpdateStatusAsync(Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[UpdateStatusResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_UpdateStatusResponse 'Trinsic.Services.VerifiableCredentials.V1.UpdateStatusResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_UpdateStatusAsync(Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## VerifiableCredential.VerifiableCredentialClient.UpdateStatusAsync(UpdateStatusRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Update credential status by setting the revocation value  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.VerifiableCredentials.V1.UpdateStatusResponse> UpdateStatusAsync(Trinsic.Services.VerifiableCredentials.V1.UpdateStatusRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_UpdateStatusAsync(Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [UpdateStatusRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest 'Trinsic.Services.VerifiableCredentials.V1.UpdateStatusRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_UpdateStatusAsync(Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_UpdateStatusAsync(Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_UpdateStatusAsync(Trinsic_Services_VerifiableCredentials_V1_UpdateStatusRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[UpdateStatusResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_UpdateStatusResponse 'Trinsic.Services.VerifiableCredentials.V1.UpdateStatusResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifyProof(Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest_Grpc_Core_CallOptions)'></a>
## VerifiableCredential.VerifiableCredentialClient.VerifyProof(VerifyProofRequest, CallOptions) Method
Verifies a proof by checking the signature value, and if possible schema validation,  
revocation status, and issuer status against a trust registry  
```csharp
public virtual Trinsic.Services.VerifiableCredentials.V1.VerifyProofResponse VerifyProof(Trinsic.Services.VerifiableCredentials.V1.VerifyProofRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifyProof(Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest_Grpc_Core_CallOptions)_request'></a>
`request` [VerifyProofRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest 'Trinsic.Services.VerifiableCredentials.V1.VerifyProofRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifyProof(Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[VerifyProofResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_VerifyProofResponse 'Trinsic.Services.VerifiableCredentials.V1.VerifyProofResponse')  
The response received from the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifyProof(Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## VerifiableCredential.VerifiableCredentialClient.VerifyProof(VerifyProofRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Verifies a proof by checking the signature value, and if possible schema validation,  
revocation status, and issuer status against a trust registry  
```csharp
public virtual Trinsic.Services.VerifiableCredentials.V1.VerifyProofResponse VerifyProof(Trinsic.Services.VerifiableCredentials.V1.VerifyProofRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifyProof(Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [VerifyProofRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest 'Trinsic.Services.VerifiableCredentials.V1.VerifyProofRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifyProof(Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifyProof(Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifyProof(Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[VerifyProofResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_VerifyProofResponse 'Trinsic.Services.VerifiableCredentials.V1.VerifyProofResponse')  
The response received from the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifyProofAsync(Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest_Grpc_Core_CallOptions)'></a>
## VerifiableCredential.VerifiableCredentialClient.VerifyProofAsync(VerifyProofRequest, CallOptions) Method
Verifies a proof by checking the signature value, and if possible schema validation,  
revocation status, and issuer status against a trust registry  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.VerifiableCredentials.V1.VerifyProofResponse> VerifyProofAsync(Trinsic.Services.VerifiableCredentials.V1.VerifyProofRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifyProofAsync(Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest_Grpc_Core_CallOptions)_request'></a>
`request` [VerifyProofRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest 'Trinsic.Services.VerifiableCredentials.V1.VerifyProofRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifyProofAsync(Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[VerifyProofResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_VerifyProofResponse 'Trinsic.Services.VerifiableCredentials.V1.VerifyProofResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifyProofAsync(Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## VerifiableCredential.VerifiableCredentialClient.VerifyProofAsync(VerifyProofRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Verifies a proof by checking the signature value, and if possible schema validation,  
revocation status, and issuer status against a trust registry  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.VerifiableCredentials.V1.VerifyProofResponse> VerifyProofAsync(Trinsic.Services.VerifiableCredentials.V1.VerifyProofRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifyProofAsync(Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [VerifyProofRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest 'Trinsic.Services.VerifiableCredentials.V1.VerifyProofRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifyProofAsync(Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifyProofAsync(Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredential_VerifiableCredentialClient_VerifyProofAsync(Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[VerifyProofResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_VerifyProofResponse 'Trinsic.Services.VerifiableCredentials.V1.VerifyProofResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredentialsReflection'></a>
## VerifiableCredentialsReflection Class
Holder for reflection information generated from services/verifiable-credentials/v1/verifiable-credentials.proto
```csharp
public static class VerifiableCredentialsReflection
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; VerifiableCredentialsReflection  
### Properties
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifiableCredentialsReflection_Descriptor'></a>
## VerifiableCredentialsReflection.Descriptor Property
File descriptor for services/verifiable-credentials/v1/verifiable-credentials.proto
```csharp
public static Google.Protobuf.Reflection.FileDescriptor Descriptor { get; }
```
#### Property Value
[Google.Protobuf.Reflection.FileDescriptor](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.Reflection.FileDescriptor 'Google.Protobuf.Reflection.FileDescriptor')
  
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest'></a>
## VerifyProofRequest Class
Verify Proof  
```csharp
public sealed class VerifyProofRequest :
Google.Protobuf.IMessage<Trinsic.Services.VerifiableCredentials.V1.VerifyProofRequest>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.VerifiableCredentials.V1.VerifyProofRequest>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.VerifiableCredentials.V1.VerifyProofRequest>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; VerifyProofRequest  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[VerifyProofRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest 'Trinsic.Services.VerifiableCredentials.V1.VerifyProofRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[VerifyProofRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest 'Trinsic.Services.VerifiableCredentials.V1.VerifyProofRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[VerifyProofRequest](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest 'Trinsic.Services.VerifiableCredentials.V1.VerifyProofRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifyProofRequest_ProofDocumentJsonFieldNumber'></a>
## VerifyProofRequest.ProofDocumentJsonFieldNumber Field
Field number for the "proof_document_json" field.
```csharp
public const int ProofDocumentJsonFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifyProofResponse'></a>
## VerifyProofResponse Class
```csharp
public sealed class VerifyProofResponse :
Google.Protobuf.IMessage<Trinsic.Services.VerifiableCredentials.V1.VerifyProofResponse>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.VerifiableCredentials.V1.VerifyProofResponse>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.VerifiableCredentials.V1.VerifyProofResponse>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; VerifyProofResponse  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[VerifyProofResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_VerifyProofResponse 'Trinsic.Services.VerifiableCredentials.V1.VerifyProofResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[VerifyProofResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_VerifyProofResponse 'Trinsic.Services.VerifiableCredentials.V1.VerifyProofResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[VerifyProofResponse](Trinsic_Services_VerifiableCredentials_V1.md#Trinsic_Services_VerifiableCredentials_V1_VerifyProofResponse 'Trinsic.Services.VerifiableCredentials.V1.VerifyProofResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifyProofResponse_IsValidFieldNumber'></a>
## VerifyProofResponse.IsValidFieldNumber Field
Field number for the "is_valid" field.
```csharp
public const int IsValidFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifyProofResponse_ValidationMessagesFieldNumber'></a>
## VerifyProofResponse.ValidationMessagesFieldNumber Field
Field number for the "validation_messages" field.
```csharp
public const int ValidationMessagesFieldNumber = 2;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
### Properties
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifyProofResponse_IsValid'></a>
## VerifyProofResponse.IsValid Property
Indicates if the proof is valid  
```csharp
public bool IsValid { get; set; }
```
#### Property Value
[System.Boolean](https://docs.microsoft.com/en-us/dotnet/api/System.Boolean 'System.Boolean')
  
<a name='Trinsic_Services_VerifiableCredentials_V1_VerifyProofResponse_ValidationMessages'></a>
## VerifyProofResponse.ValidationMessages Property
Validation messages that describe invalid verifications  
based on different factors, such as schema validation,  
proof verification, revocation registry membership, etc.  
If the proof is not valid, this field will contain detailed  
results where this verification failed.  
```csharp
public Google.Protobuf.Collections.RepeatedField<string> ValidationMessages { get; }
```
#### Property Value
[Google.Protobuf.Collections.RepeatedField&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.Collections.RepeatedField-1 'Google.Protobuf.Collections.RepeatedField`1')[System.String](https://docs.microsoft.com/en-us/dotnet/api/System.String 'System.String')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.Collections.RepeatedField-1 'Google.Protobuf.Collections.RepeatedField`1')
  
  
### Enums
<a name='Trinsic_Services_VerifiableCredentials_V1_CreateProofRequest_ProofOneofCase'></a>
## CreateProofRequest.ProofOneofCase Enum
Enum of possible cases for the "proof" oneof.
```csharp
public enum CreateProofRequest.ProofOneofCase

```
  
<a name='Trinsic_Services_VerifiableCredentials_V1_SendRequest_DeliveryMethodOneofCase'></a>
## SendRequest.DeliveryMethodOneofCase Enum
Enum of possible cases for the "delivery_method" oneof.
```csharp
public enum SendRequest.DeliveryMethodOneofCase

```
  
