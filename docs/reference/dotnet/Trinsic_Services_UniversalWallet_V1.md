#### [Trinsic](index.md 'index')
## Trinsic.Services.UniversalWallet.V1 Namespace
### Classes
<a name='Trinsic_Services_UniversalWallet_V1_DeleteItemRequest'></a>
## DeleteItemRequest Class
Delete item request  
```csharp
public sealed class DeleteItemRequest :
Google.Protobuf.IMessage<Trinsic.Services.UniversalWallet.V1.DeleteItemRequest>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.UniversalWallet.V1.DeleteItemRequest>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.UniversalWallet.V1.DeleteItemRequest>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; DeleteItemRequest  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[DeleteItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_DeleteItemRequest 'Trinsic.Services.UniversalWallet.V1.DeleteItemRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[DeleteItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_DeleteItemRequest 'Trinsic.Services.UniversalWallet.V1.DeleteItemRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[DeleteItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_DeleteItemRequest 'Trinsic.Services.UniversalWallet.V1.DeleteItemRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_UniversalWallet_V1_DeleteItemRequest_ItemIdFieldNumber'></a>
## DeleteItemRequest.ItemIdFieldNumber Field
Field number for the "item_id" field.
```csharp
public const int ItemIdFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
### Properties
<a name='Trinsic_Services_UniversalWallet_V1_DeleteItemRequest_ItemId'></a>
## DeleteItemRequest.ItemId Property
item identifier of the record to delete  
```csharp
public string ItemId { get; set; }
```
#### Property Value
[System.String](https://docs.microsoft.com/en-us/dotnet/api/System.String 'System.String')
  
  
<a name='Trinsic_Services_UniversalWallet_V1_DeleteItemResponse'></a>
## DeleteItemResponse Class
Delete item response  
```csharp
public sealed class DeleteItemResponse :
Google.Protobuf.IMessage<Trinsic.Services.UniversalWallet.V1.DeleteItemResponse>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.UniversalWallet.V1.DeleteItemResponse>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.UniversalWallet.V1.DeleteItemResponse>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; DeleteItemResponse  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[DeleteItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_DeleteItemResponse 'Trinsic.Services.UniversalWallet.V1.DeleteItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[DeleteItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_DeleteItemResponse 'Trinsic.Services.UniversalWallet.V1.DeleteItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[DeleteItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_DeleteItemResponse 'Trinsic.Services.UniversalWallet.V1.DeleteItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_UniversalWallet_V1_DeleteItemResponse_StatusFieldNumber'></a>
## DeleteItemResponse.StatusFieldNumber Field
Field number for the "status" field.
```csharp
public const int StatusFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
  
<a name='Trinsic_Services_UniversalWallet_V1_GetItemRequest'></a>
## GetItemRequest Class
Get item request object  
```csharp
public sealed class GetItemRequest :
Google.Protobuf.IMessage<Trinsic.Services.UniversalWallet.V1.GetItemRequest>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.UniversalWallet.V1.GetItemRequest>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.UniversalWallet.V1.GetItemRequest>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; GetItemRequest  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[GetItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_GetItemRequest 'Trinsic.Services.UniversalWallet.V1.GetItemRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[GetItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_GetItemRequest 'Trinsic.Services.UniversalWallet.V1.GetItemRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[GetItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_GetItemRequest 'Trinsic.Services.UniversalWallet.V1.GetItemRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_UniversalWallet_V1_GetItemRequest_ItemIdFieldNumber'></a>
## GetItemRequest.ItemIdFieldNumber Field
Field number for the "item_id" field.
```csharp
public const int ItemIdFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
### Properties
<a name='Trinsic_Services_UniversalWallet_V1_GetItemRequest_ItemId'></a>
## GetItemRequest.ItemId Property
The item identifier  
```csharp
public string ItemId { get; set; }
```
#### Property Value
[System.String](https://docs.microsoft.com/en-us/dotnet/api/System.String 'System.String')
  
  
<a name='Trinsic_Services_UniversalWallet_V1_GetItemResponse'></a>
## GetItemResponse Class
Get item response object  
```csharp
public sealed class GetItemResponse :
Google.Protobuf.IMessage<Trinsic.Services.UniversalWallet.V1.GetItemResponse>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.UniversalWallet.V1.GetItemResponse>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.UniversalWallet.V1.GetItemResponse>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; GetItemResponse  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[GetItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_GetItemResponse 'Trinsic.Services.UniversalWallet.V1.GetItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[GetItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_GetItemResponse 'Trinsic.Services.UniversalWallet.V1.GetItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[GetItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_GetItemResponse 'Trinsic.Services.UniversalWallet.V1.GetItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_UniversalWallet_V1_GetItemResponse_ItemJsonFieldNumber'></a>
## GetItemResponse.ItemJsonFieldNumber Field
Field number for the "item_json" field.
```csharp
public const int ItemJsonFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
<a name='Trinsic_Services_UniversalWallet_V1_GetItemResponse_ItemTypeFieldNumber'></a>
## GetItemResponse.ItemTypeFieldNumber Field
Field number for the "item_type" field.
```csharp
public const int ItemTypeFieldNumber = 2;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
### Properties
<a name='Trinsic_Services_UniversalWallet_V1_GetItemResponse_ItemJson'></a>
## GetItemResponse.ItemJson Property
The item data represented as stringified JSON  
```csharp
public string ItemJson { get; set; }
```
#### Property Value
[System.String](https://docs.microsoft.com/en-us/dotnet/api/System.String 'System.String')
  
<a name='Trinsic_Services_UniversalWallet_V1_GetItemResponse_ItemType'></a>
## GetItemResponse.ItemType Property
User set item type that described the content of this item  
```csharp
public string ItemType { get; set; }
```
#### Property Value
[System.String](https://docs.microsoft.com/en-us/dotnet/api/System.String 'System.String')
  
  
<a name='Trinsic_Services_UniversalWallet_V1_InsertItemRequest'></a>
## InsertItemRequest Class
Insert item request  
```csharp
public sealed class InsertItemRequest :
Google.Protobuf.IMessage<Trinsic.Services.UniversalWallet.V1.InsertItemRequest>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.UniversalWallet.V1.InsertItemRequest>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.UniversalWallet.V1.InsertItemRequest>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; InsertItemRequest  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[InsertItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_InsertItemRequest 'Trinsic.Services.UniversalWallet.V1.InsertItemRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[InsertItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_InsertItemRequest 'Trinsic.Services.UniversalWallet.V1.InsertItemRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[InsertItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_InsertItemRequest 'Trinsic.Services.UniversalWallet.V1.InsertItemRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_UniversalWallet_V1_InsertItemRequest_ItemJsonFieldNumber'></a>
## InsertItemRequest.ItemJsonFieldNumber Field
Field number for the "item_json" field.
```csharp
public const int ItemJsonFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
<a name='Trinsic_Services_UniversalWallet_V1_InsertItemRequest_ItemTypeFieldNumber'></a>
## InsertItemRequest.ItemTypeFieldNumber Field
Field number for the "item_type" field.
```csharp
public const int ItemTypeFieldNumber = 2;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
### Properties
<a name='Trinsic_Services_UniversalWallet_V1_InsertItemRequest_ItemJson'></a>
## InsertItemRequest.ItemJson Property
the document to insert as stringified json  
```csharp
public string ItemJson { get; set; }
```
#### Property Value
[System.String](https://docs.microsoft.com/en-us/dotnet/api/System.String 'System.String')
  
<a name='Trinsic_Services_UniversalWallet_V1_InsertItemRequest_ItemType'></a>
## InsertItemRequest.ItemType Property
optional item type ex. "VerifiableCredential"  
```csharp
public string ItemType { get; set; }
```
#### Property Value
[System.String](https://docs.microsoft.com/en-us/dotnet/api/System.String 'System.String')
  
  
<a name='Trinsic_Services_UniversalWallet_V1_InsertItemResponse'></a>
## InsertItemResponse Class
Insert item response  
```csharp
public sealed class InsertItemResponse :
Google.Protobuf.IMessage<Trinsic.Services.UniversalWallet.V1.InsertItemResponse>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.UniversalWallet.V1.InsertItemResponse>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.UniversalWallet.V1.InsertItemResponse>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; InsertItemResponse  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[InsertItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_InsertItemResponse 'Trinsic.Services.UniversalWallet.V1.InsertItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[InsertItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_InsertItemResponse 'Trinsic.Services.UniversalWallet.V1.InsertItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[InsertItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_InsertItemResponse 'Trinsic.Services.UniversalWallet.V1.InsertItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_UniversalWallet_V1_InsertItemResponse_ItemIdFieldNumber'></a>
## InsertItemResponse.ItemIdFieldNumber Field
Field number for the "item_id" field.
```csharp
public const int ItemIdFieldNumber = 2;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
<a name='Trinsic_Services_UniversalWallet_V1_InsertItemResponse_StatusFieldNumber'></a>
## InsertItemResponse.StatusFieldNumber Field
Field number for the "status" field.
```csharp
public const int StatusFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
### Properties
<a name='Trinsic_Services_UniversalWallet_V1_InsertItemResponse_ItemId'></a>
## InsertItemResponse.ItemId Property
The item identifier of the inserted record  
```csharp
public string ItemId { get; set; }
```
#### Property Value
[System.String](https://docs.microsoft.com/en-us/dotnet/api/System.String 'System.String')
  
  
<a name='Trinsic_Services_UniversalWallet_V1_SearchRequest'></a>
## SearchRequest Class
Search request object  
```csharp
public sealed class SearchRequest :
Google.Protobuf.IMessage<Trinsic.Services.UniversalWallet.V1.SearchRequest>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.UniversalWallet.V1.SearchRequest>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.UniversalWallet.V1.SearchRequest>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; SearchRequest  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[SearchRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_SearchRequest 'Trinsic.Services.UniversalWallet.V1.SearchRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[SearchRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_SearchRequest 'Trinsic.Services.UniversalWallet.V1.SearchRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[SearchRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_SearchRequest 'Trinsic.Services.UniversalWallet.V1.SearchRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_UniversalWallet_V1_SearchRequest_ContinuationTokenFieldNumber'></a>
## SearchRequest.ContinuationTokenFieldNumber Field
Field number for the "continuation_token" field.
```csharp
public const int ContinuationTokenFieldNumber = 2;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
<a name='Trinsic_Services_UniversalWallet_V1_SearchRequest_QueryFieldNumber'></a>
## SearchRequest.QueryFieldNumber Field
Field number for the "query" field.
```csharp
public const int QueryFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
  
<a name='Trinsic_Services_UniversalWallet_V1_SearchResponse'></a>
## SearchResponse Class
Search response object  
```csharp
public sealed class SearchResponse :
Google.Protobuf.IMessage<Trinsic.Services.UniversalWallet.V1.SearchResponse>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.UniversalWallet.V1.SearchResponse>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.UniversalWallet.V1.SearchResponse>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; SearchResponse  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[SearchResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_SearchResponse 'Trinsic.Services.UniversalWallet.V1.SearchResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[SearchResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_SearchResponse 'Trinsic.Services.UniversalWallet.V1.SearchResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[SearchResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_SearchResponse 'Trinsic.Services.UniversalWallet.V1.SearchResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_UniversalWallet_V1_SearchResponse_ContinuationTokenFieldNumber'></a>
## SearchResponse.ContinuationTokenFieldNumber Field
Field number for the "continuation_token" field.
```csharp
public const int ContinuationTokenFieldNumber = 4;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
<a name='Trinsic_Services_UniversalWallet_V1_SearchResponse_CountFieldNumber'></a>
## SearchResponse.CountFieldNumber Field
Field number for the "count" field.
```csharp
public const int CountFieldNumber = 3;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
<a name='Trinsic_Services_UniversalWallet_V1_SearchResponse_HasMoreFieldNumber'></a>
## SearchResponse.HasMoreFieldNumber Field
Field number for the "has_more" field.
```csharp
public const int HasMoreFieldNumber = 2;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
<a name='Trinsic_Services_UniversalWallet_V1_SearchResponse_ItemsFieldNumber'></a>
## SearchResponse.ItemsFieldNumber Field
Field number for the "items" field.
```csharp
public const int ItemsFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet'></a>
## UniversalWallet Class
```csharp
public static class UniversalWallet
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; UniversalWallet  
### Properties
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_Descriptor'></a>
## UniversalWallet.Descriptor Property
Service descriptor
```csharp
public static Google.Protobuf.Reflection.ServiceDescriptor Descriptor { get; }
```
#### Property Value
[Google.Protobuf.Reflection.ServiceDescriptor](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.Reflection.ServiceDescriptor 'Google.Protobuf.Reflection.ServiceDescriptor')
  
### Methods
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_BindService(Grpc_Core_ServiceBinderBase_Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase)'></a>
## UniversalWallet.BindService(ServiceBinderBase, UniversalWalletBase) Method
Register service method with a service binder with or without implementation. Useful when customizing the  service binding logic.  
            Note: this method is part of an experimental API that can change or be removed without any prior notice.
```csharp
public static void BindService(Grpc.Core.ServiceBinderBase serviceBinder, Trinsic.Services.UniversalWallet.V1.UniversalWallet.UniversalWalletBase serviceImpl);
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_BindService(Grpc_Core_ServiceBinderBase_Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase)_serviceBinder'></a>
`serviceBinder` [Grpc.Core.ServiceBinderBase](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ServiceBinderBase 'Grpc.Core.ServiceBinderBase')  
Service methods will be bound by calling `AddMethod` on this object.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_BindService(Grpc_Core_ServiceBinderBase_Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase)_serviceImpl'></a>
`serviceImpl` [UniversalWalletBase](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase 'Trinsic.Services.UniversalWallet.V1.UniversalWallet.UniversalWalletBase')  
An object implementing the server-side handling logic.
  
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_BindService(Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase)'></a>
## UniversalWallet.BindService(UniversalWalletBase) Method
Creates service definition that can be registered with a server
```csharp
public static Grpc.Core.ServerServiceDefinition BindService(Trinsic.Services.UniversalWallet.V1.UniversalWallet.UniversalWalletBase serviceImpl);
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_BindService(Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase)_serviceImpl'></a>
`serviceImpl` [UniversalWalletBase](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase 'Trinsic.Services.UniversalWallet.V1.UniversalWallet.UniversalWalletBase')  
An object implementing the server-side handling logic.
  
#### Returns
[Grpc.Core.ServerServiceDefinition](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ServerServiceDefinition 'Grpc.Core.ServerServiceDefinition')  
  
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase'></a>
## UniversalWallet.UniversalWalletBase Class
Base class for server-side implementations of UniversalWallet
```csharp
public abstract class UniversalWallet.UniversalWalletBase
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; UniversalWalletBase  
### Methods
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase_DeleteItem(Trinsic_Services_UniversalWallet_V1_DeleteItemRequest_Grpc_Core_ServerCallContext)'></a>
## UniversalWallet.UniversalWalletBase.DeleteItem(DeleteItemRequest, ServerCallContext) Method
Delete an item from the wallet permanently  
```csharp
public virtual System.Threading.Tasks.Task<Trinsic.Services.UniversalWallet.V1.DeleteItemResponse> DeleteItem(Trinsic.Services.UniversalWallet.V1.DeleteItemRequest request, Grpc.Core.ServerCallContext context);
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase_DeleteItem(Trinsic_Services_UniversalWallet_V1_DeleteItemRequest_Grpc_Core_ServerCallContext)_request'></a>
`request` [DeleteItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_DeleteItemRequest 'Trinsic.Services.UniversalWallet.V1.DeleteItemRequest')  
The request received from the client.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase_DeleteItem(Trinsic_Services_UniversalWallet_V1_DeleteItemRequest_Grpc_Core_ServerCallContext)_context'></a>
`context` [Grpc.Core.ServerCallContext](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ServerCallContext 'Grpc.Core.ServerCallContext')  
The context of the server-side call handler being invoked.
  
#### Returns
[System.Threading.Tasks.Task&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')[DeleteItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_DeleteItemResponse 'Trinsic.Services.UniversalWallet.V1.DeleteItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')  
The response to send back to the client (wrapped by a task).
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase_GetItem(Trinsic_Services_UniversalWallet_V1_GetItemRequest_Grpc_Core_ServerCallContext)'></a>
## UniversalWallet.UniversalWalletBase.GetItem(GetItemRequest, ServerCallContext) Method
Retrieve an item from the wallet with a given item identifier  
```csharp
public virtual System.Threading.Tasks.Task<Trinsic.Services.UniversalWallet.V1.GetItemResponse> GetItem(Trinsic.Services.UniversalWallet.V1.GetItemRequest request, Grpc.Core.ServerCallContext context);
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase_GetItem(Trinsic_Services_UniversalWallet_V1_GetItemRequest_Grpc_Core_ServerCallContext)_request'></a>
`request` [GetItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_GetItemRequest 'Trinsic.Services.UniversalWallet.V1.GetItemRequest')  
The request received from the client.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase_GetItem(Trinsic_Services_UniversalWallet_V1_GetItemRequest_Grpc_Core_ServerCallContext)_context'></a>
`context` [Grpc.Core.ServerCallContext](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ServerCallContext 'Grpc.Core.ServerCallContext')  
The context of the server-side call handler being invoked.
  
#### Returns
[System.Threading.Tasks.Task&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')[GetItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_GetItemResponse 'Trinsic.Services.UniversalWallet.V1.GetItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')  
The response to send back to the client (wrapped by a task).
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase_InsertItem(Trinsic_Services_UniversalWallet_V1_InsertItemRequest_Grpc_Core_ServerCallContext)'></a>
## UniversalWallet.UniversalWalletBase.InsertItem(InsertItemRequest, ServerCallContext) Method
Insert an item into the wallet  
```csharp
public virtual System.Threading.Tasks.Task<Trinsic.Services.UniversalWallet.V1.InsertItemResponse> InsertItem(Trinsic.Services.UniversalWallet.V1.InsertItemRequest request, Grpc.Core.ServerCallContext context);
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase_InsertItem(Trinsic_Services_UniversalWallet_V1_InsertItemRequest_Grpc_Core_ServerCallContext)_request'></a>
`request` [InsertItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_InsertItemRequest 'Trinsic.Services.UniversalWallet.V1.InsertItemRequest')  
The request received from the client.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase_InsertItem(Trinsic_Services_UniversalWallet_V1_InsertItemRequest_Grpc_Core_ServerCallContext)_context'></a>
`context` [Grpc.Core.ServerCallContext](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ServerCallContext 'Grpc.Core.ServerCallContext')  
The context of the server-side call handler being invoked.
  
#### Returns
[System.Threading.Tasks.Task&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')[InsertItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_InsertItemResponse 'Trinsic.Services.UniversalWallet.V1.InsertItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')  
The response to send back to the client (wrapped by a task).
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase_Search(Trinsic_Services_UniversalWallet_V1_SearchRequest_Grpc_Core_ServerCallContext)'></a>
## UniversalWallet.UniversalWalletBase.Search(SearchRequest, ServerCallContext) Method
Search the wallet using a SQL-like syntax  
```csharp
public virtual System.Threading.Tasks.Task<Trinsic.Services.UniversalWallet.V1.SearchResponse> Search(Trinsic.Services.UniversalWallet.V1.SearchRequest request, Grpc.Core.ServerCallContext context);
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase_Search(Trinsic_Services_UniversalWallet_V1_SearchRequest_Grpc_Core_ServerCallContext)_request'></a>
`request` [SearchRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_SearchRequest 'Trinsic.Services.UniversalWallet.V1.SearchRequest')  
The request received from the client.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase_Search(Trinsic_Services_UniversalWallet_V1_SearchRequest_Grpc_Core_ServerCallContext)_context'></a>
`context` [Grpc.Core.ServerCallContext](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ServerCallContext 'Grpc.Core.ServerCallContext')  
The context of the server-side call handler being invoked.
  
#### Returns
[System.Threading.Tasks.Task&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')[SearchResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_SearchResponse 'Trinsic.Services.UniversalWallet.V1.SearchResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')  
The response to send back to the client (wrapped by a task).
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase_UpdateItem(Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_Grpc_Core_ServerCallContext)'></a>
## UniversalWallet.UniversalWalletBase.UpdateItem(UpdateItemRequest, ServerCallContext) Method
Insert an item into the wallet  
```csharp
public virtual System.Threading.Tasks.Task<Trinsic.Services.UniversalWallet.V1.UpdateItemResponse> UpdateItem(Trinsic.Services.UniversalWallet.V1.UpdateItemRequest request, Grpc.Core.ServerCallContext context);
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase_UpdateItem(Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_Grpc_Core_ServerCallContext)_request'></a>
`request` [UpdateItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_UpdateItemRequest 'Trinsic.Services.UniversalWallet.V1.UpdateItemRequest')  
The request received from the client.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletBase_UpdateItem(Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_Grpc_Core_ServerCallContext)_context'></a>
`context` [Grpc.Core.ServerCallContext](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ServerCallContext 'Grpc.Core.ServerCallContext')  
The context of the server-side call handler being invoked.
  
#### Returns
[System.Threading.Tasks.Task&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')[UpdateItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_UpdateItemResponse 'Trinsic.Services.UniversalWallet.V1.UpdateItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.Tasks.Task-1 'System.Threading.Tasks.Task`1')  
The response to send back to the client (wrapped by a task).
  
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient'></a>
## UniversalWallet.UniversalWalletClient Class
Client for UniversalWallet
```csharp
public class UniversalWallet.UniversalWalletClient : Grpc.Core.ClientBase<Trinsic.Services.UniversalWallet.V1.UniversalWallet.UniversalWalletClient>
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; [Grpc.Core.ClientBase](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ClientBase 'Grpc.Core.ClientBase') &#129106; [Grpc.Core.ClientBase&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ClientBase-1 'Grpc.Core.ClientBase`1')[UniversalWalletClient](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient 'Trinsic.Services.UniversalWallet.V1.UniversalWallet.UniversalWalletClient')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ClientBase-1 'Grpc.Core.ClientBase`1') &#129106; UniversalWalletClient  
### Constructors
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UniversalWalletClient()'></a>
## UniversalWallet.UniversalWalletClient.UniversalWalletClient() Constructor
Protected parameterless constructor to allow creation of test doubles.
```csharp
protected UniversalWalletClient();
```
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UniversalWalletClient(Grpc_Core_CallInvoker)'></a>
## UniversalWallet.UniversalWalletClient.UniversalWalletClient(CallInvoker) Constructor
Creates a new client for UniversalWallet that uses a custom `CallInvoker`.
```csharp
public UniversalWalletClient(Grpc.Core.CallInvoker callInvoker);
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UniversalWalletClient(Grpc_Core_CallInvoker)_callInvoker'></a>
`callInvoker` [Grpc.Core.CallInvoker](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallInvoker 'Grpc.Core.CallInvoker')  
The callInvoker to use to make remote calls.
  
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UniversalWalletClient(Grpc_Core_ChannelBase)'></a>
## UniversalWallet.UniversalWalletClient.UniversalWalletClient(ChannelBase) Constructor
Creates a new client for UniversalWallet
```csharp
public UniversalWalletClient(Grpc.Core.ChannelBase channel);
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UniversalWalletClient(Grpc_Core_ChannelBase)_channel'></a>
`channel` [Grpc.Core.ChannelBase](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ChannelBase 'Grpc.Core.ChannelBase')  
The channel to use to make remote calls.
  
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UniversalWalletClient(Grpc_Core_ClientBase_ClientBaseConfiguration)'></a>
## UniversalWallet.UniversalWalletClient.UniversalWalletClient(ClientBaseConfiguration) Constructor
Protected constructor to allow creation of configured clients.
```csharp
protected UniversalWalletClient(Grpc.Core.ClientBase.ClientBaseConfiguration configuration);
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UniversalWalletClient(Grpc_Core_ClientBase_ClientBaseConfiguration)_configuration'></a>
`configuration` [Grpc.Core.ClientBase.ClientBaseConfiguration](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ClientBase.ClientBaseConfiguration 'Grpc.Core.ClientBase.ClientBaseConfiguration')  
The client configuration.
  
  
### Methods
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_DeleteItem(Trinsic_Services_UniversalWallet_V1_DeleteItemRequest_Grpc_Core_CallOptions)'></a>
## UniversalWallet.UniversalWalletClient.DeleteItem(DeleteItemRequest, CallOptions) Method
Delete an item from the wallet permanently  
```csharp
public virtual Trinsic.Services.UniversalWallet.V1.DeleteItemResponse DeleteItem(Trinsic.Services.UniversalWallet.V1.DeleteItemRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_DeleteItem(Trinsic_Services_UniversalWallet_V1_DeleteItemRequest_Grpc_Core_CallOptions)_request'></a>
`request` [DeleteItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_DeleteItemRequest 'Trinsic.Services.UniversalWallet.V1.DeleteItemRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_DeleteItem(Trinsic_Services_UniversalWallet_V1_DeleteItemRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[DeleteItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_DeleteItemResponse 'Trinsic.Services.UniversalWallet.V1.DeleteItemResponse')  
The response received from the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_DeleteItem(Trinsic_Services_UniversalWallet_V1_DeleteItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## UniversalWallet.UniversalWalletClient.DeleteItem(DeleteItemRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Delete an item from the wallet permanently  
```csharp
public virtual Trinsic.Services.UniversalWallet.V1.DeleteItemResponse DeleteItem(Trinsic.Services.UniversalWallet.V1.DeleteItemRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_DeleteItem(Trinsic_Services_UniversalWallet_V1_DeleteItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [DeleteItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_DeleteItemRequest 'Trinsic.Services.UniversalWallet.V1.DeleteItemRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_DeleteItem(Trinsic_Services_UniversalWallet_V1_DeleteItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_DeleteItem(Trinsic_Services_UniversalWallet_V1_DeleteItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_DeleteItem(Trinsic_Services_UniversalWallet_V1_DeleteItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[DeleteItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_DeleteItemResponse 'Trinsic.Services.UniversalWallet.V1.DeleteItemResponse')  
The response received from the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_DeleteItemAsync(Trinsic_Services_UniversalWallet_V1_DeleteItemRequest_Grpc_Core_CallOptions)'></a>
## UniversalWallet.UniversalWalletClient.DeleteItemAsync(DeleteItemRequest, CallOptions) Method
Delete an item from the wallet permanently  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.UniversalWallet.V1.DeleteItemResponse> DeleteItemAsync(Trinsic.Services.UniversalWallet.V1.DeleteItemRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_DeleteItemAsync(Trinsic_Services_UniversalWallet_V1_DeleteItemRequest_Grpc_Core_CallOptions)_request'></a>
`request` [DeleteItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_DeleteItemRequest 'Trinsic.Services.UniversalWallet.V1.DeleteItemRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_DeleteItemAsync(Trinsic_Services_UniversalWallet_V1_DeleteItemRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[DeleteItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_DeleteItemResponse 'Trinsic.Services.UniversalWallet.V1.DeleteItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_DeleteItemAsync(Trinsic_Services_UniversalWallet_V1_DeleteItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## UniversalWallet.UniversalWalletClient.DeleteItemAsync(DeleteItemRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Delete an item from the wallet permanently  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.UniversalWallet.V1.DeleteItemResponse> DeleteItemAsync(Trinsic.Services.UniversalWallet.V1.DeleteItemRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_DeleteItemAsync(Trinsic_Services_UniversalWallet_V1_DeleteItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [DeleteItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_DeleteItemRequest 'Trinsic.Services.UniversalWallet.V1.DeleteItemRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_DeleteItemAsync(Trinsic_Services_UniversalWallet_V1_DeleteItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_DeleteItemAsync(Trinsic_Services_UniversalWallet_V1_DeleteItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_DeleteItemAsync(Trinsic_Services_UniversalWallet_V1_DeleteItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[DeleteItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_DeleteItemResponse 'Trinsic.Services.UniversalWallet.V1.DeleteItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_GetItem(Trinsic_Services_UniversalWallet_V1_GetItemRequest_Grpc_Core_CallOptions)'></a>
## UniversalWallet.UniversalWalletClient.GetItem(GetItemRequest, CallOptions) Method
Retrieve an item from the wallet with a given item identifier  
```csharp
public virtual Trinsic.Services.UniversalWallet.V1.GetItemResponse GetItem(Trinsic.Services.UniversalWallet.V1.GetItemRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_GetItem(Trinsic_Services_UniversalWallet_V1_GetItemRequest_Grpc_Core_CallOptions)_request'></a>
`request` [GetItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_GetItemRequest 'Trinsic.Services.UniversalWallet.V1.GetItemRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_GetItem(Trinsic_Services_UniversalWallet_V1_GetItemRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[GetItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_GetItemResponse 'Trinsic.Services.UniversalWallet.V1.GetItemResponse')  
The response received from the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_GetItem(Trinsic_Services_UniversalWallet_V1_GetItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## UniversalWallet.UniversalWalletClient.GetItem(GetItemRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Retrieve an item from the wallet with a given item identifier  
```csharp
public virtual Trinsic.Services.UniversalWallet.V1.GetItemResponse GetItem(Trinsic.Services.UniversalWallet.V1.GetItemRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_GetItem(Trinsic_Services_UniversalWallet_V1_GetItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [GetItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_GetItemRequest 'Trinsic.Services.UniversalWallet.V1.GetItemRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_GetItem(Trinsic_Services_UniversalWallet_V1_GetItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_GetItem(Trinsic_Services_UniversalWallet_V1_GetItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_GetItem(Trinsic_Services_UniversalWallet_V1_GetItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[GetItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_GetItemResponse 'Trinsic.Services.UniversalWallet.V1.GetItemResponse')  
The response received from the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_GetItemAsync(Trinsic_Services_UniversalWallet_V1_GetItemRequest_Grpc_Core_CallOptions)'></a>
## UniversalWallet.UniversalWalletClient.GetItemAsync(GetItemRequest, CallOptions) Method
Retrieve an item from the wallet with a given item identifier  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.UniversalWallet.V1.GetItemResponse> GetItemAsync(Trinsic.Services.UniversalWallet.V1.GetItemRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_GetItemAsync(Trinsic_Services_UniversalWallet_V1_GetItemRequest_Grpc_Core_CallOptions)_request'></a>
`request` [GetItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_GetItemRequest 'Trinsic.Services.UniversalWallet.V1.GetItemRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_GetItemAsync(Trinsic_Services_UniversalWallet_V1_GetItemRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[GetItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_GetItemResponse 'Trinsic.Services.UniversalWallet.V1.GetItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_GetItemAsync(Trinsic_Services_UniversalWallet_V1_GetItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## UniversalWallet.UniversalWalletClient.GetItemAsync(GetItemRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Retrieve an item from the wallet with a given item identifier  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.UniversalWallet.V1.GetItemResponse> GetItemAsync(Trinsic.Services.UniversalWallet.V1.GetItemRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_GetItemAsync(Trinsic_Services_UniversalWallet_V1_GetItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [GetItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_GetItemRequest 'Trinsic.Services.UniversalWallet.V1.GetItemRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_GetItemAsync(Trinsic_Services_UniversalWallet_V1_GetItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_GetItemAsync(Trinsic_Services_UniversalWallet_V1_GetItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_GetItemAsync(Trinsic_Services_UniversalWallet_V1_GetItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[GetItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_GetItemResponse 'Trinsic.Services.UniversalWallet.V1.GetItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_InsertItem(Trinsic_Services_UniversalWallet_V1_InsertItemRequest_Grpc_Core_CallOptions)'></a>
## UniversalWallet.UniversalWalletClient.InsertItem(InsertItemRequest, CallOptions) Method
Insert an item into the wallet  
```csharp
public virtual Trinsic.Services.UniversalWallet.V1.InsertItemResponse InsertItem(Trinsic.Services.UniversalWallet.V1.InsertItemRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_InsertItem(Trinsic_Services_UniversalWallet_V1_InsertItemRequest_Grpc_Core_CallOptions)_request'></a>
`request` [InsertItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_InsertItemRequest 'Trinsic.Services.UniversalWallet.V1.InsertItemRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_InsertItem(Trinsic_Services_UniversalWallet_V1_InsertItemRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[InsertItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_InsertItemResponse 'Trinsic.Services.UniversalWallet.V1.InsertItemResponse')  
The response received from the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_InsertItem(Trinsic_Services_UniversalWallet_V1_InsertItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## UniversalWallet.UniversalWalletClient.InsertItem(InsertItemRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Insert an item into the wallet  
```csharp
public virtual Trinsic.Services.UniversalWallet.V1.InsertItemResponse InsertItem(Trinsic.Services.UniversalWallet.V1.InsertItemRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_InsertItem(Trinsic_Services_UniversalWallet_V1_InsertItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [InsertItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_InsertItemRequest 'Trinsic.Services.UniversalWallet.V1.InsertItemRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_InsertItem(Trinsic_Services_UniversalWallet_V1_InsertItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_InsertItem(Trinsic_Services_UniversalWallet_V1_InsertItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_InsertItem(Trinsic_Services_UniversalWallet_V1_InsertItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[InsertItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_InsertItemResponse 'Trinsic.Services.UniversalWallet.V1.InsertItemResponse')  
The response received from the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_InsertItemAsync(Trinsic_Services_UniversalWallet_V1_InsertItemRequest_Grpc_Core_CallOptions)'></a>
## UniversalWallet.UniversalWalletClient.InsertItemAsync(InsertItemRequest, CallOptions) Method
Insert an item into the wallet  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.UniversalWallet.V1.InsertItemResponse> InsertItemAsync(Trinsic.Services.UniversalWallet.V1.InsertItemRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_InsertItemAsync(Trinsic_Services_UniversalWallet_V1_InsertItemRequest_Grpc_Core_CallOptions)_request'></a>
`request` [InsertItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_InsertItemRequest 'Trinsic.Services.UniversalWallet.V1.InsertItemRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_InsertItemAsync(Trinsic_Services_UniversalWallet_V1_InsertItemRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[InsertItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_InsertItemResponse 'Trinsic.Services.UniversalWallet.V1.InsertItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_InsertItemAsync(Trinsic_Services_UniversalWallet_V1_InsertItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## UniversalWallet.UniversalWalletClient.InsertItemAsync(InsertItemRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Insert an item into the wallet  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.UniversalWallet.V1.InsertItemResponse> InsertItemAsync(Trinsic.Services.UniversalWallet.V1.InsertItemRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_InsertItemAsync(Trinsic_Services_UniversalWallet_V1_InsertItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [InsertItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_InsertItemRequest 'Trinsic.Services.UniversalWallet.V1.InsertItemRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_InsertItemAsync(Trinsic_Services_UniversalWallet_V1_InsertItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_InsertItemAsync(Trinsic_Services_UniversalWallet_V1_InsertItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_InsertItemAsync(Trinsic_Services_UniversalWallet_V1_InsertItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[InsertItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_InsertItemResponse 'Trinsic.Services.UniversalWallet.V1.InsertItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_NewInstance(Grpc_Core_ClientBase_ClientBaseConfiguration)'></a>
## UniversalWallet.UniversalWalletClient.NewInstance(ClientBaseConfiguration) Method
Creates a new instance of client from given `ClientBaseConfiguration`.
```csharp
protected override Trinsic.Services.UniversalWallet.V1.UniversalWallet.UniversalWalletClient NewInstance(Grpc.Core.ClientBase.ClientBaseConfiguration configuration);
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_NewInstance(Grpc_Core_ClientBase_ClientBaseConfiguration)_configuration'></a>
`configuration` [Grpc.Core.ClientBase.ClientBaseConfiguration](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.ClientBase.ClientBaseConfiguration 'Grpc.Core.ClientBase.ClientBaseConfiguration')  
  
#### Returns
[UniversalWalletClient](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient 'Trinsic.Services.UniversalWallet.V1.UniversalWallet.UniversalWalletClient')  
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_Search(Trinsic_Services_UniversalWallet_V1_SearchRequest_Grpc_Core_CallOptions)'></a>
## UniversalWallet.UniversalWalletClient.Search(SearchRequest, CallOptions) Method
Search the wallet using a SQL-like syntax  
```csharp
public virtual Trinsic.Services.UniversalWallet.V1.SearchResponse Search(Trinsic.Services.UniversalWallet.V1.SearchRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_Search(Trinsic_Services_UniversalWallet_V1_SearchRequest_Grpc_Core_CallOptions)_request'></a>
`request` [SearchRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_SearchRequest 'Trinsic.Services.UniversalWallet.V1.SearchRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_Search(Trinsic_Services_UniversalWallet_V1_SearchRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[SearchResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_SearchResponse 'Trinsic.Services.UniversalWallet.V1.SearchResponse')  
The response received from the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_Search(Trinsic_Services_UniversalWallet_V1_SearchRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## UniversalWallet.UniversalWalletClient.Search(SearchRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Search the wallet using a SQL-like syntax  
```csharp
public virtual Trinsic.Services.UniversalWallet.V1.SearchResponse Search(Trinsic.Services.UniversalWallet.V1.SearchRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_Search(Trinsic_Services_UniversalWallet_V1_SearchRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [SearchRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_SearchRequest 'Trinsic.Services.UniversalWallet.V1.SearchRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_Search(Trinsic_Services_UniversalWallet_V1_SearchRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_Search(Trinsic_Services_UniversalWallet_V1_SearchRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_Search(Trinsic_Services_UniversalWallet_V1_SearchRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[SearchResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_SearchResponse 'Trinsic.Services.UniversalWallet.V1.SearchResponse')  
The response received from the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_SearchAsync(Trinsic_Services_UniversalWallet_V1_SearchRequest_Grpc_Core_CallOptions)'></a>
## UniversalWallet.UniversalWalletClient.SearchAsync(SearchRequest, CallOptions) Method
Search the wallet using a SQL-like syntax  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.UniversalWallet.V1.SearchResponse> SearchAsync(Trinsic.Services.UniversalWallet.V1.SearchRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_SearchAsync(Trinsic_Services_UniversalWallet_V1_SearchRequest_Grpc_Core_CallOptions)_request'></a>
`request` [SearchRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_SearchRequest 'Trinsic.Services.UniversalWallet.V1.SearchRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_SearchAsync(Trinsic_Services_UniversalWallet_V1_SearchRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[SearchResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_SearchResponse 'Trinsic.Services.UniversalWallet.V1.SearchResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_SearchAsync(Trinsic_Services_UniversalWallet_V1_SearchRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## UniversalWallet.UniversalWalletClient.SearchAsync(SearchRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Search the wallet using a SQL-like syntax  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.UniversalWallet.V1.SearchResponse> SearchAsync(Trinsic.Services.UniversalWallet.V1.SearchRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_SearchAsync(Trinsic_Services_UniversalWallet_V1_SearchRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [SearchRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_SearchRequest 'Trinsic.Services.UniversalWallet.V1.SearchRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_SearchAsync(Trinsic_Services_UniversalWallet_V1_SearchRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_SearchAsync(Trinsic_Services_UniversalWallet_V1_SearchRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_SearchAsync(Trinsic_Services_UniversalWallet_V1_SearchRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[SearchResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_SearchResponse 'Trinsic.Services.UniversalWallet.V1.SearchResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UpdateItem(Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_Grpc_Core_CallOptions)'></a>
## UniversalWallet.UniversalWalletClient.UpdateItem(UpdateItemRequest, CallOptions) Method
Insert an item into the wallet  
```csharp
public virtual Trinsic.Services.UniversalWallet.V1.UpdateItemResponse UpdateItem(Trinsic.Services.UniversalWallet.V1.UpdateItemRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UpdateItem(Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_Grpc_Core_CallOptions)_request'></a>
`request` [UpdateItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_UpdateItemRequest 'Trinsic.Services.UniversalWallet.V1.UpdateItemRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UpdateItem(Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[UpdateItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_UpdateItemResponse 'Trinsic.Services.UniversalWallet.V1.UpdateItemResponse')  
The response received from the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UpdateItem(Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## UniversalWallet.UniversalWalletClient.UpdateItem(UpdateItemRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Insert an item into the wallet  
```csharp
public virtual Trinsic.Services.UniversalWallet.V1.UpdateItemResponse UpdateItem(Trinsic.Services.UniversalWallet.V1.UpdateItemRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UpdateItem(Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [UpdateItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_UpdateItemRequest 'Trinsic.Services.UniversalWallet.V1.UpdateItemRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UpdateItem(Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UpdateItem(Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UpdateItem(Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[UpdateItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_UpdateItemResponse 'Trinsic.Services.UniversalWallet.V1.UpdateItemResponse')  
The response received from the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UpdateItemAsync(Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_Grpc_Core_CallOptions)'></a>
## UniversalWallet.UniversalWalletClient.UpdateItemAsync(UpdateItemRequest, CallOptions) Method
Insert an item into the wallet  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.UniversalWallet.V1.UpdateItemResponse> UpdateItemAsync(Trinsic.Services.UniversalWallet.V1.UpdateItemRequest request, Grpc.Core.CallOptions options);
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UpdateItemAsync(Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_Grpc_Core_CallOptions)_request'></a>
`request` [UpdateItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_UpdateItemRequest 'Trinsic.Services.UniversalWallet.V1.UpdateItemRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UpdateItemAsync(Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_Grpc_Core_CallOptions)_options'></a>
`options` [Grpc.Core.CallOptions](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.CallOptions 'Grpc.Core.CallOptions')  
The options for the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[UpdateItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_UpdateItemResponse 'Trinsic.Services.UniversalWallet.V1.UpdateItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UpdateItemAsync(Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)'></a>
## UniversalWallet.UniversalWalletClient.UpdateItemAsync(UpdateItemRequest, Metadata, Nullable&lt;DateTime&gt;, CancellationToken) Method
Insert an item into the wallet  
```csharp
public virtual Grpc.Core.AsyncUnaryCall<Trinsic.Services.UniversalWallet.V1.UpdateItemResponse> UpdateItemAsync(Trinsic.Services.UniversalWallet.V1.UpdateItemRequest request, Grpc.Core.Metadata headers=null, System.Nullable<System.DateTime> deadline=null, System.Threading.CancellationToken cancellationToken=default(System.Threading.CancellationToken));
```
#### Parameters
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UpdateItemAsync(Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_request'></a>
`request` [UpdateItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_UpdateItemRequest 'Trinsic.Services.UniversalWallet.V1.UpdateItemRequest')  
The request to send to the server.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UpdateItemAsync(Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_headers'></a>
`headers` [Grpc.Core.Metadata](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.Metadata 'Grpc.Core.Metadata')  
The initial metadata to send with the call. This parameter is optional.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UpdateItemAsync(Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_deadline'></a>
`deadline` [System.Nullable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')[System.DateTime](https://docs.microsoft.com/en-us/dotnet/api/System.DateTime 'System.DateTime')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.Nullable-1 'System.Nullable`1')  
An optional deadline for the call. The call will be cancelled if deadline is hit.
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWallet_UniversalWalletClient_UpdateItemAsync(Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_Grpc_Core_Metadata_System_Nullable_System_DateTime__System_Threading_CancellationToken)_cancellationToken'></a>
`cancellationToken` [System.Threading.CancellationToken](https://docs.microsoft.com/en-us/dotnet/api/System.Threading.CancellationToken 'System.Threading.CancellationToken')  
An optional token for canceling the call.
  
#### Returns
[Grpc.Core.AsyncUnaryCall&lt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')[UpdateItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_UpdateItemResponse 'Trinsic.Services.UniversalWallet.V1.UpdateItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Grpc.Core.AsyncUnaryCall-1 'Grpc.Core.AsyncUnaryCall`1')  
The call object.
  
  
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWalletReflection'></a>
## UniversalWalletReflection Class
Holder for reflection information generated from services/universal-wallet/v1/universal-wallet.proto
```csharp
public static class UniversalWalletReflection
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; UniversalWalletReflection  
### Properties
<a name='Trinsic_Services_UniversalWallet_V1_UniversalWalletReflection_Descriptor'></a>
## UniversalWalletReflection.Descriptor Property
File descriptor for services/universal-wallet/v1/universal-wallet.proto
```csharp
public static Google.Protobuf.Reflection.FileDescriptor Descriptor { get; }
```
#### Property Value
[Google.Protobuf.Reflection.FileDescriptor](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.Reflection.FileDescriptor 'Google.Protobuf.Reflection.FileDescriptor')
  
  
<a name='Trinsic_Services_UniversalWallet_V1_UpdateItemRequest'></a>
## UpdateItemRequest Class
Update item request object  
```csharp
public sealed class UpdateItemRequest :
Google.Protobuf.IMessage<Trinsic.Services.UniversalWallet.V1.UpdateItemRequest>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.UniversalWallet.V1.UpdateItemRequest>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.UniversalWallet.V1.UpdateItemRequest>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; UpdateItemRequest  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[UpdateItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_UpdateItemRequest 'Trinsic.Services.UniversalWallet.V1.UpdateItemRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[UpdateItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_UpdateItemRequest 'Trinsic.Services.UniversalWallet.V1.UpdateItemRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[UpdateItemRequest](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_UpdateItemRequest 'Trinsic.Services.UniversalWallet.V1.UpdateItemRequest')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_ItemIdFieldNumber'></a>
## UpdateItemRequest.ItemIdFieldNumber Field
Field number for the "item_id" field.
```csharp
public const int ItemIdFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
<a name='Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_ItemTypeFieldNumber'></a>
## UpdateItemRequest.ItemTypeFieldNumber Field
Field number for the "item_type" field.
```csharp
public const int ItemTypeFieldNumber = 2;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
### Properties
<a name='Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_ItemId'></a>
## UpdateItemRequest.ItemId Property
The item identifier  
```csharp
public string ItemId { get; set; }
```
#### Property Value
[System.String](https://docs.microsoft.com/en-us/dotnet/api/System.String 'System.String')
  
<a name='Trinsic_Services_UniversalWallet_V1_UpdateItemRequest_ItemType'></a>
## UpdateItemRequest.ItemType Property
The item type that described the content of this item  
```csharp
public string ItemType { get; set; }
```
#### Property Value
[System.String](https://docs.microsoft.com/en-us/dotnet/api/System.String 'System.String')
  
  
<a name='Trinsic_Services_UniversalWallet_V1_UpdateItemResponse'></a>
## UpdateItemResponse Class
Update item response object  
```csharp
public sealed class UpdateItemResponse :
Google.Protobuf.IMessage<Trinsic.Services.UniversalWallet.V1.UpdateItemResponse>,
Google.Protobuf.IMessage,
System.IEquatable<Trinsic.Services.UniversalWallet.V1.UpdateItemResponse>,
Google.Protobuf.IDeepCloneable<Trinsic.Services.UniversalWallet.V1.UpdateItemResponse>,
Google.Protobuf.IBufferMessage
```

Inheritance [System.Object](https://docs.microsoft.com/en-us/dotnet/api/System.Object 'System.Object') &#129106; UpdateItemResponse  

Implements [Google.Protobuf.IMessage&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1')[UpdateItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_UpdateItemResponse 'Trinsic.Services.UniversalWallet.V1.UpdateItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage-1 'Google.Protobuf.IMessage`1'), [Google.Protobuf.IMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IMessage 'Google.Protobuf.IMessage'), [System.IEquatable&lt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1')[UpdateItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_UpdateItemResponse 'Trinsic.Services.UniversalWallet.V1.UpdateItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/System.IEquatable-1 'System.IEquatable`1'), [Google.Protobuf.IDeepCloneable&lt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1')[UpdateItemResponse](Trinsic_Services_UniversalWallet_V1.md#Trinsic_Services_UniversalWallet_V1_UpdateItemResponse 'Trinsic.Services.UniversalWallet.V1.UpdateItemResponse')[&gt;](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IDeepCloneable-1 'Google.Protobuf.IDeepCloneable`1'), [Google.Protobuf.IBufferMessage](https://docs.microsoft.com/en-us/dotnet/api/Google.Protobuf.IBufferMessage 'Google.Protobuf.IBufferMessage')  
### Fields
<a name='Trinsic_Services_UniversalWallet_V1_UpdateItemResponse_StatusFieldNumber'></a>
## UpdateItemResponse.StatusFieldNumber Field
Field number for the "status" field.
```csharp
public const int StatusFieldNumber = 1;
```
#### Field Value
[System.Int32](https://docs.microsoft.com/en-us/dotnet/api/System.Int32 'System.Int32')
  
### Properties
<a name='Trinsic_Services_UniversalWallet_V1_UpdateItemResponse_Status'></a>
## UpdateItemResponse.Status Property
Response status  
```csharp
public Trinsic.Services.Common.V1.ResponseStatus Status { get; set; }
```
#### Property Value
[Trinsic.Services.Common.V1.ResponseStatus](https://docs.microsoft.com/en-us/dotnet/api/Trinsic.Services.Common.V1.ResponseStatus 'Trinsic.Services.Common.V1.ResponseStatus')
  
  
