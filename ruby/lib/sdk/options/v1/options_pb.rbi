# Code generated by protoc-gen-rbi. DO NOT EDIT.
# source: sdk/options/v1/options.proto
# typed: strict

module Sdk; end
module Sdk::Options; end
module Sdk::Options::V1; end

class Sdk::Options::V1::ServiceOptions
  include Google::Protobuf
  include Google::Protobuf::MessageExts
  extend Google::Protobuf::MessageExts::ClassMethods

  sig { params(str: String).returns(Sdk::Options::V1::ServiceOptions) }
  def self.decode(str)
  end

  sig { params(msg: Sdk::Options::V1::ServiceOptions).returns(String) }
  def self.encode(msg)
  end

  sig { params(str: String, kw: T.untyped).returns(Sdk::Options::V1::ServiceOptions) }
  def self.decode_json(str, **kw)
  end

  sig { params(msg: Sdk::Options::V1::ServiceOptions, kw: T.untyped).returns(String) }
  def self.encode_json(msg, **kw)
  end

  sig { returns(Google::Protobuf::Descriptor) }
  def self.descriptor
  end

  sig do
    params(
      server: T.nilable(Sdk::Options::V1::ServerConfiguration),
      profile: T.nilable(Services::Account::V1::AccountProfile),
      ecosystem: T.nilable(String)
    ).void
  end
  def initialize(
    server: nil,
    profile: nil,
    ecosystem: ""
  )
  end

  sig { returns(T.nilable(Sdk::Options::V1::ServerConfiguration)) }
  def server
  end

  sig { params(value: T.nilable(Sdk::Options::V1::ServerConfiguration)).void }
  def server=(value)
  end

  sig { void }
  def clear_server
  end

  sig { returns(T.nilable(Services::Account::V1::AccountProfile)) }
  def profile
  end

  sig { params(value: T.nilable(Services::Account::V1::AccountProfile)).void }
  def profile=(value)
  end

  sig { void }
  def clear_profile
  end

  sig { returns(String) }
  def ecosystem
  end

  sig { params(value: String).void }
  def ecosystem=(value)
  end

  sig { void }
  def clear_ecosystem
  end

  sig { params(field: String).returns(T.untyped) }
  def [](field)
  end

  sig { params(field: String, value: T.untyped).void }
  def []=(field, value)
  end

  sig { returns(T::Hash[Symbol, T.untyped]) }
  def to_h
  end
end

class Sdk::Options::V1::ServerConfiguration
  include Google::Protobuf
  include Google::Protobuf::MessageExts
  extend Google::Protobuf::MessageExts::ClassMethods

  sig { params(str: String).returns(Sdk::Options::V1::ServerConfiguration) }
  def self.decode(str)
  end

  sig { params(msg: Sdk::Options::V1::ServerConfiguration).returns(String) }
  def self.encode(msg)
  end

  sig { params(str: String, kw: T.untyped).returns(Sdk::Options::V1::ServerConfiguration) }
  def self.decode_json(str, **kw)
  end

  sig { params(msg: Sdk::Options::V1::ServerConfiguration, kw: T.untyped).returns(String) }
  def self.encode_json(msg, **kw)
  end

  sig { returns(Google::Protobuf::Descriptor) }
  def self.descriptor
  end

  sig do
    params(
      endpoint: T.nilable(String),
      port: T.nilable(Integer),
      use_tls: T.nilable(T::Boolean)
    ).void
  end
  def initialize(
    endpoint: "",
    port: 0,
    use_tls: false
  )
  end

  sig { returns(String) }
  def endpoint
  end

  sig { params(value: String).void }
  def endpoint=(value)
  end

  sig { void }
  def clear_endpoint
  end

  sig { returns(Integer) }
  def port
  end

  sig { params(value: Integer).void }
  def port=(value)
  end

  sig { void }
  def clear_port
  end

  sig { returns(T::Boolean) }
  def use_tls
  end

  sig { params(value: T::Boolean).void }
  def use_tls=(value)
  end

  sig { void }
  def clear_use_tls
  end

  sig { params(field: String).returns(T.untyped) }
  def [](field)
  end

  sig { params(field: String, value: T.untyped).void }
  def []=(field, value)
  end

  sig { returns(T::Hash[Symbol, T.untyped]) }
  def to_h
  end
end