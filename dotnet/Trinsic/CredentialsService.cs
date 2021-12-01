using System;
using System.Threading.Tasks;
using Google.Protobuf.WellKnownTypes;
using Newtonsoft.Json.Linq;
using Trinsic.Services.Account.V1;
using Trinsic.Services.Common.V1;
using Trinsic.Services.VerifiableCredentials.V1;

namespace Trinsic;

public class CredentialsService : ServiceBase
{
    public CredentialsService(AccountProfile accountProfile, ServerConfig? serverConfig = null)
        : base(accountProfile, serverConfig)
    {
        Client = new VerifiableCredential.VerifiableCredentialClient(Channel);
    }

    internal VerifiableCredential.VerifiableCredentialClient Client { get; }

    /// <summary>
    /// Signs an input credential
    /// </summary>
    /// <param name="document"></param>
    /// <returns></returns>
    public async Task<JObject> IssueCredential(JObject document)
    {
        try
        {
            IssueRequest request = new()
            {
                Document = new JsonPayload { JsonStruct = document.ToStruct() }
            };
            var response = await Client.IssueAsync(request, await BuildMetadataAsync(request));
            return response.Document.JsonStruct.ToJObject();
        }
        catch (Exception e)
        {
            Console.WriteLine(e);
            throw;
        }
    }


    /// <summary>
    /// Derive a proof from an existing document in the wallet using
    /// an input reveal document frame
    /// </summary>
    /// <param name="documentId"></param>
    /// <param name="revealDocument"></param>
    /// <returns></returns>
    public async Task<JObject> CreateProof(string documentId, JObject revealDocument)
    {
        CreateProofRequest request = new()
        {
            DocumentId = documentId,
            RevealDocument = new JsonPayload { JsonStruct = revealDocument.ToStruct() }
        };
        var response = await Client.CreateProofAsync(
            request: request,
            headers: await BuildMetadataAsync(request));

        return response.ProofDocument.JsonStruct.ToJObject();
    }

    /// <summary>
    /// Verifies a proof document
    /// </summary>
    /// <param name="proofDocument"></param>
    /// <returns></returns>
    public async Task<bool> VerifyProof(JObject proofDocument)
    {
        VerifyProofRequest request = new()
        {
            ProofDocument = new JsonPayload { JsonString = proofDocument.ToString() }
        };
        var response = await Client.VerifyProofAsync(
            request: request,
            headers: await BuildMetadataAsync(request));

        return response.Valid;
    }


    /// <summary>
    /// Sends a document to the specified destination
    /// </summary>
    /// <param name="document"></param>
    /// <param name="email"></param>
    /// <returns></returns>
    public async Task Send(JObject document, string email)
    {
        SendRequest request = new()
        {
            Email = email,
            Document = new JsonPayload { JsonStruct = document.ToStruct() }
        };
        var response = await Client.SendAsync(
            request: request,
            headers: await BuildMetadataAsync(request));
    }
}