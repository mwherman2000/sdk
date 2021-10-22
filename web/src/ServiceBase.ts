import { CreateProofRequest, JsonWebKey, LdProofs, LdSuite } from "@trinsic/okapi";
import { Metadata } from "grpc-web";
import { WalletProfile } from "./proto";
import { Struct } from "google-protobuf/google/protobuf/struct_pb";
import { Buffer } from "buffer";

export default abstract class ServiceBase {
  capInvocation: string;

  getMetadata(): Metadata {
    if (!this.capInvocation) throw new Error("Profile not set.");
    let metadata = { "Capability-Invocation": this.capInvocation };
    return metadata;
  }

  async setProfile(profile: WalletProfile): Promise<void> {
    let capabilityDocument = {
      "@context": "https://wid.org/security/v2",
      invocationTarget: profile.getWalletId(),
      proof: {
        proofPurpose: "capabilityInvocation",
        created: new Date().toISOString(),
        capability: profile.getCapability(),
      },
    };

    let proofRequest = new CreateProofRequest()
      .setDocument(Struct.fromJavaScript(capabilityDocument))
      .setKey(JsonWebKey.deserializeBinary(profile.getInvokerJwk_asU8()))
      .setSuite(LdSuite.JCSED25519SIGNATURE2020);

    let proofResponse = await LdProofs.generate(proofRequest);

    // Set the auth field to the signed document by converting it back
    // to JSON and encoding it in base64
    this.capInvocation = Buffer.from(JSON.stringify(proofResponse.getSignedDocument().toJavaScript())).toString(
      "base64"
    );
  }
}
