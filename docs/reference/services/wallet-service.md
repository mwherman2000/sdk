# Wallet Service

The wallet service is the main interface for interacting with a cloud wallet. The service endpoints are designed to closely match the recommendations of the [Universal Wallet 2020 <small>:material-open-in-new:<small>](https://w3c-ccg.github.io/universal-wallet-interop-spec/){target=_blank} specification by W3C Community Credentials Group. The service exposes a gRPC interface and a set of data contracts as described in the specification. Our intention with this design is to bring it closer to interoperability as more implementations of this wallet appear in production.


## Create Wallet

Wallets can be created directly by the user or through an invitation by the ecosystem provider. Depending on the ecosystem settings, direct wallet creation may not be enabled for your provider.

### Create wallet directly

To create a wallet directly without an invitation, use the following methods. These methods return secure profile data that should be stored in a safe place. The profile is used to authenticate the wallet service with the cloud provider. Read more details and recommendations in the [Security Profiles](/reference/#authorization) page.

=== "Trinsic CLI"
    ```bash
    trinsic wallet create --name <profile_name>
    ```
=== "TypeScript"
    ```typescript
    const profile = await walletService.createWallet();
    ```
=== "C#"
    ```csharp
    var profile = await walletService.Create();
    ```
=== "Python"
    ```python 
    profile = await wallet_service.create_wallet()
    ```

### Create wallet with provider invitation

If invited by a provider, you can supply the security code found in your invitation (via email, SMS, etc). Read more about [inviting participants](/reference/services/provider-service/#invite-participants) to your ecosystem as provider.

=== "Trinsic CLI"
    ```bash
    trinsic wallet create --name <profile> --security-code <code>
    ```
=== "TypeScript"
    ```typescript
    const profile = await walletService.createWallet("<security code>");
    ```
=== "C#"
    ```csharp
    var profile = await walletService.Create("<security code>");
    ```
=== "Python"
    ```python
    profile = await wallet_service.create_wallet("<security code>")
    ```

## Insert Item

This method allows inserting any JSON data in the wallet.

=== "Trinsic CLI"
    ```bash
    trinsic wallet insert-item --item <INPUT_JSON_FILE>
    ```
=== "TypeScript"
    ```typescript
    let itemId = await walletService.insertItem({
        "foo": "bar"
    });
    ```
=== "C#"
    ```csharp
    var item = new JObject
    {
        { "foo", "bar" }
    };

    var itemId = await walletService.InsertItem(item);
    ```
=== "Python"
    ```python
    item_id = await wallet_service.insert_item(credential)
    ```
The output of this method will be a unique `itemId` that can be used as input where required.

## Search / Query

Querying wallet data in our SDK is enabled through the use of familiar SQL syntax. All data is stored in JSON-LD format, so it can be easily searched.
This apporach allows us to give developers full control over how data is retrieved. In addition to customizable sorting, paging and filtering, developers have the ability to construct projections, combine result sets, and even run user-defined functions over their queries.

> This endpoint will support querying using [Verifiable Presentation Request Spec <small>:material-open-in-new:<small>](https://w3c-ccg.github.io/vp-request-spec/){target=_blank}. This feature is still in development.

### Basic Search

The default query used in the commands below returns a full wallet result set. The query is `SELECT * FROM c`.

=== "Trinsic CLI"
    ```bash
    trinsic wallet search
    ```
=== "TypeScript"

    ```typescript
    const items = await walletService.Search();
    ```

=== "C#"

    ```csharp
    var items = await walletService.Search();
    ```

=== "Python"
    ```python
    item = await wallet_service.search()
    ```

### SQL Search

To pass custom query to the search function, use the query parameter or the available overload.

=== "Trinsic CLI"
    ```bash
    trinsic wallet search \
        --query "SELECT * FROM c WHERE c.type = 'VerifiableCredential'"
    ```
=== "TypeScript"

    ```typescript
    const query = "SELECT * FROM c WHERE c.type = 'VerifiableCredential'";

    const items = await walletService.search(query);
    ```

=== "C#"

    ```csharp
    var items = await walletService.Search("SELECT * FROM c WHERE c.type = 'VerifiableCredential'");
    ```

=== "Python"
    ```python
    query = "SELECT * FROM c WHERE c.type = 'VerifiableCredential'"
    item = await wallet_service.search(query)
    ```


### Common SQL Queries

#### Paging

Paging uses the `OFFSET` clause that takes in a value indicating how many records should be skipped in the returned query. To specify the size of the result set (page size) use the `LIMIT` clause.

```sql
SELECT * FROM c OFFSET 10 LIMIT 5
```

#### Sorting

The optional `ORDER BY` clause specifies the sorting order for results returned by the query. To control sorting order, specify `ASC` or `DESC` at the end; if not specified ascending order is used by default.

```sql
SELECT * FROM c ORDER BY c.credential.issued DESC
```

#### Filtering

The optional WHERE clause (`WHERE <filter_condition>`) specifies condition(s) that the source JSON items must satisfy for the query to include them in results. A JSON item must evaluate the specified conditions to true to be considered for the result. The index layer uses the `WHERE` clause to determine the smallest subset of source items that can be part of the result.

```sql
SELECT * FROM c WHERE c.name = 'Trinsic' AND c.dateCreated >= "2020-09-30T23:14:25.7251173Z"
```

#### Grouping

The `GROUP BY` clause divides the query's results according to the values of one or more specified properties.
Examples and detailed description on working with grouped results [can be found here <small>:material-open-in-new:<small>](https://docs.microsoft.com/en-us/azure/cosmos-db/sql-query-group-by){target=_blank}

### Additional Resources

You can read the full documentation on working with SQL queries on the [Azure Cosmos DB website <small>:material-open-in-new:<small>](https://docs.microsoft.com/en-us/azure/cosmos-db/sql-query-getting-started){target=_blank}.

