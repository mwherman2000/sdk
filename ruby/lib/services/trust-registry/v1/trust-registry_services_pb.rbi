# Code generated by protoc-gen-rbi. DO NOT EDIT.
# source: services/trust-registry/v1/trust-registry.proto
# typed: strict

module Services::Trustregistry::V1::TrustRegistry
  class Service
    include GRPC::GenericService
  end

  class Stub < GRPC::ClientStub
    sig do
      params(
        host: String,
        creds: T.any(GRPC::Core::ChannelCredentials, Symbol),
        kw: T.untyped,
      ).void
    end
    def initialize(host, creds, **kw)
    end

    sig do
      params(
        request: Services::Trustregistry::V1::AddFrameworkRequest
      ).returns(Services::Trustregistry::V1::AddFrameworkResponse)
    end
    def add_framework(request)
    end

    sig do
      params(
        request: Services::Trustregistry::V1::RemoveFrameworkRequest
      ).returns(Services::Trustregistry::V1::RemoveFrameworkResponse)
    end
    def remove_framework(request)
    end

    sig do
      params(
        request: Services::Trustregistry::V1::SearchRegistryRequest
      ).returns(Services::Trustregistry::V1::SearchRegistryResponse)
    end
    def search_registry(request)
    end

    sig do
      params(
        request: Services::Trustregistry::V1::RegisterIssuerRequest
      ).returns(Services::Trustregistry::V1::RegisterIssuerResponse)
    end
    def register_issuer(request)
    end

    sig do
      params(
        request: Services::Trustregistry::V1::RegisterVerifierRequest
      ).returns(Services::Trustregistry::V1::RegisterVerifierResponse)
    end
    def register_verifier(request)
    end

    sig do
      params(
        request: Services::Trustregistry::V1::UnregisterIssuerRequest
      ).returns(Services::Trustregistry::V1::UnregisterIssuerResponse)
    end
    def unregister_issuer(request)
    end

    sig do
      params(
        request: Services::Trustregistry::V1::UnregisterVerifierRequest
      ).returns(Services::Trustregistry::V1::UnregisterVerifierResponse)
    end
    def unregister_verifier(request)
    end

    sig do
      params(
        request: Services::Trustregistry::V1::CheckIssuerStatusRequest
      ).returns(Services::Trustregistry::V1::CheckIssuerStatusResponse)
    end
    def check_issuer_status(request)
    end

    sig do
      params(
        request: Services::Trustregistry::V1::CheckVerifierStatusRequest
      ).returns(Services::Trustregistry::V1::CheckVerifierStatusResponse)
    end
    def check_verifier_status(request)
    end

    sig do
      params(
        request: Services::Trustregistry::V1::FetchDataRequest
      ).returns(T::Enumerable[Services::Trustregistry::V1::FetchDataResponse])
    end
    def fetch_data(request)
    end
  end
end
