# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: CoreService.proto, DebugService.proto, IssuerService.proto, ProviderService.proto, WalletService.proto
# plugin: python-betterproto
from dataclasses import dataclass
from typing import Dict, List, Optional

import betterproto
from betterproto.grpc.grpclib_server import ServiceBase
import grpclib


class ResponseStatus(betterproto.Enum):
    SUCCESS = 0
    WALLET_ACCESS_DENIED = 10
    WALLET_EXISTS = 11
    ITEM_NOT_FOUND = 20
    SERIALIZATION_ERROR = 200
    UNKNOWN_ERROR = 100


class JsonFormat(betterproto.Enum):
    Protobuf = 0
    Binary = 1
    String = 2


class ParticipantType(betterproto.Enum):
    participant_type_individual = 0
    participant_type_organization = 1


class InvitationStatusResponseStatus(betterproto.Enum):
    Error = 0
    InvitationSent = 1
    Completed = 2


@dataclass(eq=False, repr=False)
class JsonPayload(betterproto.Message):
    json_struct: "betterproto_lib_google_protobuf.Struct" = betterproto.message_field(
        1, group="json"
    )
    json_string: str = betterproto.string_field(2, group="json")
    json_bytes: bytes = betterproto.bytes_field(3, group="json")


@dataclass(eq=False, repr=False)
class IssueRequest(betterproto.Message):
    document: "JsonPayload" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class IssueResponse(betterproto.Message):
    document: "JsonPayload" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class CreateProofRequest(betterproto.Message):
    """Create Proof"""

    reveal_document: "JsonPayload" = betterproto.message_field(1)
    document_id: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class CreateProofResponse(betterproto.Message):
    proof_document: "JsonPayload" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class VerifyProofRequest(betterproto.Message):
    """Verify Proof"""

    proof_document: "JsonPayload" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class VerifyProofResponse(betterproto.Message):
    valid: bool = betterproto.bool_field(1)


@dataclass(eq=False, repr=False)
class SendRequest(betterproto.Message):
    email: str = betterproto.string_field(1, group="delivery_method")
    did_uri: str = betterproto.string_field(2, group="delivery_method")
    didcomm_invitation: "JsonPayload" = betterproto.message_field(
        3, group="delivery_method"
    )
    document: "JsonPayload" = betterproto.message_field(100)


@dataclass(eq=False, repr=False)
class SendResponse(betterproto.Message):
    status: "ResponseStatus" = betterproto.enum_field(1)


@dataclass(eq=False, repr=False)
class InviteRequest(betterproto.Message):
    participant: "ParticipantType" = betterproto.enum_field(1)
    description: str = betterproto.string_field(2)
    email: str = betterproto.string_field(5, group="contact_method")
    phone: str = betterproto.string_field(6, group="contact_method")
    didcomm_invitation: "InviteRequestDidCommInvitation" = betterproto.message_field(
        7, group="contact_method"
    )


@dataclass(eq=False, repr=False)
class InviteRequestDidCommInvitation(betterproto.Message):
    pass


@dataclass(eq=False, repr=False)
class InviteResponse(betterproto.Message):
    status: "ResponseStatus" = betterproto.enum_field(1)
    invitation_id: str = betterproto.string_field(10)


@dataclass(eq=False, repr=False)
class InvitationStatusRequest(betterproto.Message):
    """
    Request details for the status of onboarding an individual or organization.
    The referenece_id passed is the response from the `Onboard` method call
    """

    invitation_id: str = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class InvitationStatusResponse(betterproto.Message):
    status: "InvitationStatusResponseStatus" = betterproto.enum_field(1)
    status_details: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class CreateWalletRequest(betterproto.Message):
    controller: str = betterproto.string_field(1)
    description: str = betterproto.string_field(2)
    # (Optional) Supply an invitation id to associate this caller profile to an
    # existing cloud wallet.
    security_code: str = betterproto.string_field(3)


@dataclass(eq=False, repr=False)
class CreateWalletResponse(betterproto.Message):
    status: "ResponseStatus" = betterproto.enum_field(1)
    wallet_id: str = betterproto.string_field(2)
    capability: str = betterproto.string_field(3)
    invoker: str = betterproto.string_field(4)


@dataclass(eq=False, repr=False)
class ConnectRequest(betterproto.Message):
    email: str = betterproto.string_field(5, group="contact_method")
    phone: str = betterproto.string_field(6, group="contact_method")


@dataclass(eq=False, repr=False)
class ConnectResponse(betterproto.Message):
    status: "ResponseStatus" = betterproto.enum_field(1)


@dataclass(eq=False, repr=False)
class InvitationToken(betterproto.Message):
    security_code: str = betterproto.string_field(1)
    wallet_id: str = betterproto.string_field(2)
    email: str = betterproto.string_field(5, group="contact_method")
    phone: str = betterproto.string_field(6, group="contact_method")


@dataclass(eq=False, repr=False)
class WalletProfile(betterproto.Message):
    """
    Stores profile data for accessing a wallet.This result should be stored
    somewhere safe,as it contains private key information.
    """

    did_document: "JsonPayload" = betterproto.message_field(1)
    wallet_id: str = betterproto.string_field(2)
    invoker: str = betterproto.string_field(3)
    capability: str = betterproto.string_field(4)
    invoker_jwk: bytes = betterproto.bytes_field(5)


@dataclass(eq=False, repr=False)
class GrantAccessRequest(betterproto.Message):
    wallet_id: str = betterproto.string_field(1)
    did: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class GrantAccessResponse(betterproto.Message):
    status: "ResponseStatus" = betterproto.enum_field(1)


@dataclass(eq=False, repr=False)
class RevokeAccessRequest(betterproto.Message):
    wallet_id: str = betterproto.string_field(1)
    did: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class RevokeAccessResponse(betterproto.Message):
    status: "ResponseStatus" = betterproto.enum_field(1)


@dataclass(eq=False, repr=False)
class GetProviderConfigurationResponse(betterproto.Message):
    did_document: "JsonPayload" = betterproto.message_field(1)
    key_agreement_key_id: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class SearchRequest(betterproto.Message):
    query: str = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class SearchResponse(betterproto.Message):
    items: List["JsonPayload"] = betterproto.message_field(1)
    has_more: bool = betterproto.bool_field(2)


@dataclass(eq=False, repr=False)
class InsertItemRequest(betterproto.Message):
    item: "JsonPayload" = betterproto.message_field(1)
    item_type: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class InsertItemResponse(betterproto.Message):
    status: "ResponseStatus" = betterproto.enum_field(1)
    item_id: str = betterproto.string_field(2)


class CommonStub(betterproto.ServiceStub):
    async def request(
        self,
        *,
        iv: bytes = b"",
        aad: bytes = b"",
        ciphertext: bytes = b"",
        tag: bytes = b"",
        recipients: Optional[List["EncryptionRecipient"]] = None,
    ) -> "__pbmse__.EncryptedMessage":
        recipients = recipients or []

        request = __pbmse__.EncryptedMessage()
        request.iv = iv
        request.aad = aad
        request.ciphertext = ciphertext
        request.tag = tag
        if recipients is not None:
            request.recipients = recipients

        return await self._unary_unary(
            "/trinsic.services.Common/Request", request, __pbmse__.EncryptedMessage
        )


class DebuggingStub(betterproto.ServiceStub):
    async def call_empty(self) -> "betterproto_lib_google_protobuf.Empty":

        request = betterproto_lib_google_protobuf.Empty()

        return await self._unary_unary(
            "/trinsic.services.Debugging/CallEmpty",
            request,
            betterproto_lib_google_protobuf.Empty,
        )


class CredentialStub(betterproto.ServiceStub):
    async def issue(self, *, document: "JsonPayload" = None) -> "IssueResponse":

        request = IssueRequest()
        if document is not None:
            request.document = document

        return await self._unary_unary(
            "/trinsic.services.Credential/Issue", request, IssueResponse
        )

    async def create_proof(
        self, *, reveal_document: "JsonPayload" = None, document_id: str = ""
    ) -> "CreateProofResponse":

        request = CreateProofRequest()
        if reveal_document is not None:
            request.reveal_document = reveal_document
        request.document_id = document_id

        return await self._unary_unary(
            "/trinsic.services.Credential/CreateProof", request, CreateProofResponse
        )

    async def verify_proof(
        self, *, proof_document: "JsonPayload" = None
    ) -> "VerifyProofResponse":

        request = VerifyProofRequest()
        if proof_document is not None:
            request.proof_document = proof_document

        return await self._unary_unary(
            "/trinsic.services.Credential/VerifyProof", request, VerifyProofResponse
        )

    async def send(
        self,
        *,
        email: str = "",
        did_uri: str = "",
        didcomm_invitation: "JsonPayload" = None,
        document: "JsonPayload" = None,
    ) -> "SendResponse":

        request = SendRequest()
        request.email = email
        request.did_uri = did_uri
        if didcomm_invitation is not None:
            request.didcomm_invitation = didcomm_invitation
        if document is not None:
            request.document = document

        return await self._unary_unary(
            "/trinsic.services.Credential/Send", request, SendResponse
        )


class ProviderStub(betterproto.ServiceStub):
    async def invite(
        self,
        *,
        participant: "ParticipantType" = None,
        description: str = "",
        email: str = "",
        phone: str = "",
        didcomm_invitation: "InviteRequestDidCommInvitation" = None,
    ) -> "InviteResponse":

        request = InviteRequest()
        request.participant = participant
        request.description = description
        request.email = email
        request.phone = phone
        if didcomm_invitation is not None:
            request.didcomm_invitation = didcomm_invitation

        return await self._unary_unary(
            "/trinsic.services.Provider/Invite", request, InviteResponse
        )

    async def invite_with_workflow(
        self,
        *,
        participant: "ParticipantType" = None,
        description: str = "",
        email: str = "",
        phone: str = "",
        didcomm_invitation: "InviteRequestDidCommInvitation" = None,
    ) -> "InviteResponse":

        request = InviteRequest()
        request.participant = participant
        request.description = description
        request.email = email
        request.phone = phone
        if didcomm_invitation is not None:
            request.didcomm_invitation = didcomm_invitation

        return await self._unary_unary(
            "/trinsic.services.Provider/InviteWithWorkflow", request, InviteResponse
        )

    async def invitation_status(
        self, *, invitation_id: str = ""
    ) -> "InvitationStatusResponse":

        request = InvitationStatusRequest()
        request.invitation_id = invitation_id

        return await self._unary_unary(
            "/trinsic.services.Provider/InvitationStatus",
            request,
            InvitationStatusResponse,
        )


class WalletStub(betterproto.ServiceStub):
    async def get_provider_configuration(self) -> "GetProviderConfigurationResponse":

        request = betterproto_lib_google_protobuf.Empty()

        return await self._unary_unary(
            "/trinsic.services.Wallet/GetProviderConfiguration",
            request,
            GetProviderConfigurationResponse,
        )

    async def connect_external_identity(
        self, *, email: str = "", phone: str = ""
    ) -> "ConnectResponse":

        request = ConnectRequest()
        request.email = email
        request.phone = phone

        return await self._unary_unary(
            "/trinsic.services.Wallet/ConnectExternalIdentity", request, ConnectResponse
        )

    async def create_wallet(
        self, *, controller: str = "", description: str = "", security_code: str = ""
    ) -> "CreateWalletResponse":

        request = CreateWalletRequest()
        request.controller = controller
        request.description = description
        request.security_code = security_code

        return await self._unary_unary(
            "/trinsic.services.Wallet/CreateWallet", request, CreateWalletResponse
        )

    async def create_wallet_with_workflow(
        self, *, controller: str = "", description: str = "", security_code: str = ""
    ) -> "CreateWalletResponse":

        request = CreateWalletRequest()
        request.controller = controller
        request.description = description
        request.security_code = security_code

        return await self._unary_unary(
            "/trinsic.services.Wallet/CreateWalletWithWorkflow",
            request,
            CreateWalletResponse,
        )

    async def create_wallet_encrypted(
        self,
        *,
        iv: bytes = b"",
        aad: bytes = b"",
        ciphertext: bytes = b"",
        tag: bytes = b"",
        recipients: Optional[List["EncryptionRecipient"]] = None,
    ) -> "__pbmse__.EncryptedMessage":
        recipients = recipients or []

        request = __pbmse__.EncryptedMessage()
        request.iv = iv
        request.aad = aad
        request.ciphertext = ciphertext
        request.tag = tag
        if recipients is not None:
            request.recipients = recipients

        return await self._unary_unary(
            "/trinsic.services.Wallet/CreateWalletEncrypted",
            request,
            __pbmse__.EncryptedMessage,
        )

    async def search(self, *, query: str = "") -> "SearchResponse":

        request = SearchRequest()
        request.query = query

        return await self._unary_unary(
            "/trinsic.services.Wallet/Search", request, SearchResponse
        )

    async def insert_item(
        self, *, item: "JsonPayload" = None, item_type: str = ""
    ) -> "InsertItemResponse":

        request = InsertItemRequest()
        if item is not None:
            request.item = item
        request.item_type = item_type

        return await self._unary_unary(
            "/trinsic.services.Wallet/InsertItem", request, InsertItemResponse
        )

    async def grant_access(
        self, *, wallet_id: str = "", did: str = ""
    ) -> "GrantAccessResponse":

        request = GrantAccessRequest()
        request.wallet_id = wallet_id
        request.did = did

        return await self._unary_unary(
            "/trinsic.services.Wallet/GrantAccess", request, GrantAccessResponse
        )

    async def revoke_access(
        self, *, wallet_id: str = "", did: str = ""
    ) -> "RevokeAccessResponse":

        request = RevokeAccessRequest()
        request.wallet_id = wallet_id
        request.did = did

        return await self._unary_unary(
            "/trinsic.services.Wallet/RevokeAccess", request, RevokeAccessResponse
        )


class CommonBase(ServiceBase):
    async def request(
        self,
        iv: bytes,
        aad: bytes,
        ciphertext: bytes,
        tag: bytes,
        recipients: Optional[List["EncryptionRecipient"]],
    ) -> "__pbmse__.EncryptedMessage":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_request(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "iv": request.iv,
            "aad": request.aad,
            "ciphertext": request.ciphertext,
            "tag": request.tag,
            "recipients": request.recipients,
        }

        response = await self.request(**request_kwargs)
        await stream.send_message(response)

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/trinsic.services.Common/Request": grpclib.const.Handler(
                self.__rpc_request,
                grpclib.const.Cardinality.UNARY_UNARY,
                __pbmse__.EncryptedMessage,
                __pbmse__.EncryptedMessage,
            ),
        }


class DebuggingBase(ServiceBase):
    async def call_empty(self) -> "betterproto_lib_google_protobuf.Empty":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_call_empty(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {}

        response = await self.call_empty(**request_kwargs)
        await stream.send_message(response)

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/trinsic.services.Debugging/CallEmpty": grpclib.const.Handler(
                self.__rpc_call_empty,
                grpclib.const.Cardinality.UNARY_UNARY,
                betterproto_lib_google_protobuf.Empty,
                betterproto_lib_google_protobuf.Empty,
            ),
        }


class CredentialBase(ServiceBase):
    async def issue(self, document: "JsonPayload") -> "IssueResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def create_proof(
        self, reveal_document: "JsonPayload", document_id: str
    ) -> "CreateProofResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def verify_proof(
        self, proof_document: "JsonPayload"
    ) -> "VerifyProofResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def send(
        self,
        email: str,
        did_uri: str,
        didcomm_invitation: "JsonPayload",
        document: "JsonPayload",
    ) -> "SendResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_issue(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "document": request.document,
        }

        response = await self.issue(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_create_proof(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "reveal_document": request.reveal_document,
            "document_id": request.document_id,
        }

        response = await self.create_proof(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_verify_proof(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "proof_document": request.proof_document,
        }

        response = await self.verify_proof(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_send(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "email": request.email,
            "did_uri": request.did_uri,
            "didcomm_invitation": request.didcomm_invitation,
            "document": request.document,
        }

        response = await self.send(**request_kwargs)
        await stream.send_message(response)

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/trinsic.services.Credential/Issue": grpclib.const.Handler(
                self.__rpc_issue,
                grpclib.const.Cardinality.UNARY_UNARY,
                IssueRequest,
                IssueResponse,
            ),
            "/trinsic.services.Credential/CreateProof": grpclib.const.Handler(
                self.__rpc_create_proof,
                grpclib.const.Cardinality.UNARY_UNARY,
                CreateProofRequest,
                CreateProofResponse,
            ),
            "/trinsic.services.Credential/VerifyProof": grpclib.const.Handler(
                self.__rpc_verify_proof,
                grpclib.const.Cardinality.UNARY_UNARY,
                VerifyProofRequest,
                VerifyProofResponse,
            ),
            "/trinsic.services.Credential/Send": grpclib.const.Handler(
                self.__rpc_send,
                grpclib.const.Cardinality.UNARY_UNARY,
                SendRequest,
                SendResponse,
            ),
        }


class ProviderBase(ServiceBase):
    async def invite(
        self,
        participant: "ParticipantType",
        description: str,
        email: str,
        phone: str,
        didcomm_invitation: "InviteRequestDidCommInvitation",
    ) -> "InviteResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def invite_with_workflow(
        self,
        participant: "ParticipantType",
        description: str,
        email: str,
        phone: str,
        didcomm_invitation: "InviteRequestDidCommInvitation",
    ) -> "InviteResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def invitation_status(self, invitation_id: str) -> "InvitationStatusResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_invite(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "participant": request.participant,
            "description": request.description,
            "email": request.email,
            "phone": request.phone,
            "didcomm_invitation": request.didcomm_invitation,
        }

        response = await self.invite(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_invite_with_workflow(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "participant": request.participant,
            "description": request.description,
            "email": request.email,
            "phone": request.phone,
            "didcomm_invitation": request.didcomm_invitation,
        }

        response = await self.invite_with_workflow(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_invitation_status(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "invitation_id": request.invitation_id,
        }

        response = await self.invitation_status(**request_kwargs)
        await stream.send_message(response)

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/trinsic.services.Provider/Invite": grpclib.const.Handler(
                self.__rpc_invite,
                grpclib.const.Cardinality.UNARY_UNARY,
                InviteRequest,
                InviteResponse,
            ),
            "/trinsic.services.Provider/InviteWithWorkflow": grpclib.const.Handler(
                self.__rpc_invite_with_workflow,
                grpclib.const.Cardinality.UNARY_UNARY,
                InviteRequest,
                InviteResponse,
            ),
            "/trinsic.services.Provider/InvitationStatus": grpclib.const.Handler(
                self.__rpc_invitation_status,
                grpclib.const.Cardinality.UNARY_UNARY,
                InvitationStatusRequest,
                InvitationStatusResponse,
            ),
        }


class WalletBase(ServiceBase):
    async def get_provider_configuration(self) -> "GetProviderConfigurationResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def connect_external_identity(
        self, email: str, phone: str
    ) -> "ConnectResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def create_wallet(
        self, controller: str, description: str, security_code: str
    ) -> "CreateWalletResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def create_wallet_with_workflow(
        self, controller: str, description: str, security_code: str
    ) -> "CreateWalletResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def create_wallet_encrypted(
        self,
        iv: bytes,
        aad: bytes,
        ciphertext: bytes,
        tag: bytes,
        recipients: Optional[List["EncryptionRecipient"]],
    ) -> "__pbmse__.EncryptedMessage":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def search(self, query: str) -> "SearchResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def insert_item(
        self, item: "JsonPayload", item_type: str
    ) -> "InsertItemResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def grant_access(self, wallet_id: str, did: str) -> "GrantAccessResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def revoke_access(self, wallet_id: str, did: str) -> "RevokeAccessResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_get_provider_configuration(
        self, stream: grpclib.server.Stream
    ) -> None:
        request = await stream.recv_message()

        request_kwargs = {}

        response = await self.get_provider_configuration(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_connect_external_identity(
        self, stream: grpclib.server.Stream
    ) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "email": request.email,
            "phone": request.phone,
        }

        response = await self.connect_external_identity(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_create_wallet(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "controller": request.controller,
            "description": request.description,
            "security_code": request.security_code,
        }

        response = await self.create_wallet(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_create_wallet_with_workflow(
        self, stream: grpclib.server.Stream
    ) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "controller": request.controller,
            "description": request.description,
            "security_code": request.security_code,
        }

        response = await self.create_wallet_with_workflow(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_create_wallet_encrypted(
        self, stream: grpclib.server.Stream
    ) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "iv": request.iv,
            "aad": request.aad,
            "ciphertext": request.ciphertext,
            "tag": request.tag,
            "recipients": request.recipients,
        }

        response = await self.create_wallet_encrypted(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_search(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "query": request.query,
        }

        response = await self.search(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_insert_item(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "item": request.item,
            "item_type": request.item_type,
        }

        response = await self.insert_item(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_grant_access(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "wallet_id": request.wallet_id,
            "did": request.did,
        }

        response = await self.grant_access(**request_kwargs)
        await stream.send_message(response)

    async def __rpc_revoke_access(self, stream: grpclib.server.Stream) -> None:
        request = await stream.recv_message()

        request_kwargs = {
            "wallet_id": request.wallet_id,
            "did": request.did,
        }

        response = await self.revoke_access(**request_kwargs)
        await stream.send_message(response)

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/trinsic.services.Wallet/GetProviderConfiguration": grpclib.const.Handler(
                self.__rpc_get_provider_configuration,
                grpclib.const.Cardinality.UNARY_UNARY,
                betterproto_lib_google_protobuf.Empty,
                GetProviderConfigurationResponse,
            ),
            "/trinsic.services.Wallet/ConnectExternalIdentity": grpclib.const.Handler(
                self.__rpc_connect_external_identity,
                grpclib.const.Cardinality.UNARY_UNARY,
                ConnectRequest,
                ConnectResponse,
            ),
            "/trinsic.services.Wallet/CreateWallet": grpclib.const.Handler(
                self.__rpc_create_wallet,
                grpclib.const.Cardinality.UNARY_UNARY,
                CreateWalletRequest,
                CreateWalletResponse,
            ),
            "/trinsic.services.Wallet/CreateWalletWithWorkflow": grpclib.const.Handler(
                self.__rpc_create_wallet_with_workflow,
                grpclib.const.Cardinality.UNARY_UNARY,
                CreateWalletRequest,
                CreateWalletResponse,
            ),
            "/trinsic.services.Wallet/CreateWalletEncrypted": grpclib.const.Handler(
                self.__rpc_create_wallet_encrypted,
                grpclib.const.Cardinality.UNARY_UNARY,
                __pbmse__.EncryptedMessage,
                __pbmse__.EncryptedMessage,
            ),
            "/trinsic.services.Wallet/Search": grpclib.const.Handler(
                self.__rpc_search,
                grpclib.const.Cardinality.UNARY_UNARY,
                SearchRequest,
                SearchResponse,
            ),
            "/trinsic.services.Wallet/InsertItem": grpclib.const.Handler(
                self.__rpc_insert_item,
                grpclib.const.Cardinality.UNARY_UNARY,
                InsertItemRequest,
                InsertItemResponse,
            ),
            "/trinsic.services.Wallet/GrantAccess": grpclib.const.Handler(
                self.__rpc_grant_access,
                grpclib.const.Cardinality.UNARY_UNARY,
                GrantAccessRequest,
                GrantAccessResponse,
            ),
            "/trinsic.services.Wallet/RevokeAccess": grpclib.const.Handler(
                self.__rpc_revoke_access,
                grpclib.const.Cardinality.UNARY_UNARY,
                RevokeAccessRequest,
                RevokeAccessResponse,
            ),
        }


from ... import pbmse as __pbmse__
import betterproto.lib.google.protobuf as betterproto_lib_google_protobuf