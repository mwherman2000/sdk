// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the Swift generator plugin for the protocol buffer compiler.
// Source: ProviderService.proto
//
// For information on using the generated types, please see the documentation:
//   https://github.com/apple/swift-protobuf/

import Foundation
import SwiftProtobuf

// If the compiler emits an error on this type, it is because this file
// was generated by a version of the `protoc` Swift plug-in that is
// incompatible with the version of SwiftProtobuf to which you are linking.
// Please ensure that you are building against the same version of the API
// that was used to generate this file.
fileprivate struct _GeneratedWithProtocGenSwiftVersion: SwiftProtobuf.ProtobufAPIVersionCheck {
  struct _2: SwiftProtobuf.ProtobufAPIVersion_2 {}
  typealias Version = _2
}

public enum Trinsic_Services_ParticipantType: SwiftProtobuf.Enum {
  public typealias RawValue = Int
  case individual // = 0
  case organization // = 1
  case UNRECOGNIZED(Int)

  public init() {
    self = .individual
  }

  public init?(rawValue: Int) {
    switch rawValue {
    case 0: self = .individual
    case 1: self = .organization
    default: self = .UNRECOGNIZED(rawValue)
    }
  }

  public var rawValue: Int {
    switch self {
    case .individual: return 0
    case .organization: return 1
    case .UNRECOGNIZED(let i): return i
    }
  }

}

#if swift(>=4.2)

extension Trinsic_Services_ParticipantType: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  public static var allCases: [Trinsic_Services_ParticipantType] = [
    .individual,
    .organization,
  ]
}

#endif  // swift(>=4.2)

public struct Trinsic_Services_InviteRequest {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var participant: Trinsic_Services_ParticipantType = .individual

  public var description_p: String = String()

  public var contactMethod: Trinsic_Services_InviteRequest.OneOf_ContactMethod? = nil

  public var email: String {
    get {
      if case .email(let v)? = contactMethod {return v}
      return String()
    }
    set {contactMethod = .email(newValue)}
  }

  public var phone: String {
    get {
      if case .phone(let v)? = contactMethod {return v}
      return String()
    }
    set {contactMethod = .phone(newValue)}
  }

  public var didcommInvitation: Trinsic_Services_InviteRequest.DidCommInvitation {
    get {
      if case .didcommInvitation(let v)? = contactMethod {return v}
      return Trinsic_Services_InviteRequest.DidCommInvitation()
    }
    set {contactMethod = .didcommInvitation(newValue)}
  }

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public enum OneOf_ContactMethod: Equatable {
    case email(String)
    case phone(String)
    case didcommInvitation(Trinsic_Services_InviteRequest.DidCommInvitation)

  #if !swift(>=4.1)
    public static func ==(lhs: Trinsic_Services_InviteRequest.OneOf_ContactMethod, rhs: Trinsic_Services_InviteRequest.OneOf_ContactMethod) -> Bool {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch (lhs, rhs) {
      case (.email, .email): return {
        guard case .email(let l) = lhs, case .email(let r) = rhs else { preconditionFailure() }
        return l == r
      }()
      case (.phone, .phone): return {
        guard case .phone(let l) = lhs, case .phone(let r) = rhs else { preconditionFailure() }
        return l == r
      }()
      case (.didcommInvitation, .didcommInvitation): return {
        guard case .didcommInvitation(let l) = lhs, case .didcommInvitation(let r) = rhs else { preconditionFailure() }
        return l == r
      }()
      default: return false
      }
    }
  #endif
  }

  public struct DidCommInvitation {
    // SwiftProtobuf.Message conformance is added in an extension below. See the
    // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
    // methods supported on all messages.

    public var unknownFields = SwiftProtobuf.UnknownStorage()

    public init() {}
  }

  public init() {}
}

public struct Trinsic_Services_InviteResponse {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var status: Trinsic_Services_ResponseStatus = .success

  public var invitationID: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

/// Request details for the status of onboarding
/// an individual or organization.
/// The referenece_id passed is the response from the
/// `Onboard` method call
public struct Trinsic_Services_InvitationStatusRequest {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var invitationID: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public init() {}
}

public struct Trinsic_Services_InvitationStatusResponse {
  // SwiftProtobuf.Message conformance is added in an extension below. See the
  // `Message` and `Message+*Additions` files in the SwiftProtobuf library for
  // methods supported on all messages.

  public var status: Trinsic_Services_InvitationStatusResponse.Status = .error

  public var statusDetails: String = String()

  public var unknownFields = SwiftProtobuf.UnknownStorage()

  public enum Status: SwiftProtobuf.Enum {
    public typealias RawValue = Int

    /// Onboarding resulted in error
    case error // = 0

    /// The participant has been invited
    case invitationSent // = 1

    /// The participant has been onboarded
    case completed // = 2
    case UNRECOGNIZED(Int)

    public init() {
      self = .error
    }

    public init?(rawValue: Int) {
      switch rawValue {
      case 0: self = .error
      case 1: self = .invitationSent
      case 2: self = .completed
      default: self = .UNRECOGNIZED(rawValue)
      }
    }

    public var rawValue: Int {
      switch self {
      case .error: return 0
      case .invitationSent: return 1
      case .completed: return 2
      case .UNRECOGNIZED(let i): return i
      }
    }

  }

  public init() {}
}

#if swift(>=4.2)

extension Trinsic_Services_InvitationStatusResponse.Status: CaseIterable {
  // The compiler won't synthesize support with the UNRECOGNIZED case.
  public static var allCases: [Trinsic_Services_InvitationStatusResponse.Status] = [
    .error,
    .invitationSent,
    .completed,
  ]
}

#endif  // swift(>=4.2)

// MARK: - Code below here is support for the SwiftProtobuf runtime.

fileprivate let _protobuf_package = "trinsic.services"

extension Trinsic_Services_ParticipantType: SwiftProtobuf._ProtoNameProviding {
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "participant_type_individual"),
    1: .same(proto: "participant_type_organization"),
  ]
}

extension Trinsic_Services_InviteRequest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".InviteRequest"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "participant"),
    2: .same(proto: "description"),
    5: .same(proto: "email"),
    6: .same(proto: "phone"),
    7: .standard(proto: "didcomm_invitation"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularEnumField(value: &self.participant) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.description_p) }()
      case 5: try {
        var v: String?
        try decoder.decodeSingularStringField(value: &v)
        if let v = v {
          if self.contactMethod != nil {try decoder.handleConflictingOneOf()}
          self.contactMethod = .email(v)
        }
      }()
      case 6: try {
        var v: String?
        try decoder.decodeSingularStringField(value: &v)
        if let v = v {
          if self.contactMethod != nil {try decoder.handleConflictingOneOf()}
          self.contactMethod = .phone(v)
        }
      }()
      case 7: try {
        var v: Trinsic_Services_InviteRequest.DidCommInvitation?
        var hadOneofValue = false
        if let current = self.contactMethod {
          hadOneofValue = true
          if case .didcommInvitation(let m) = current {v = m}
        }
        try decoder.decodeSingularMessageField(value: &v)
        if let v = v {
          if hadOneofValue {try decoder.handleConflictingOneOf()}
          self.contactMethod = .didcommInvitation(v)
        }
      }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if self.participant != .individual {
      try visitor.visitSingularEnumField(value: self.participant, fieldNumber: 1)
    }
    if !self.description_p.isEmpty {
      try visitor.visitSingularStringField(value: self.description_p, fieldNumber: 2)
    }
    // The use of inline closures is to circumvent an issue where the compiler
    // allocates stack space for every case branch when no optimizations are
    // enabled. https://github.com/apple/swift-protobuf/issues/1034
    switch self.contactMethod {
    case .email?: try {
      guard case .email(let v)? = self.contactMethod else { preconditionFailure() }
      try visitor.visitSingularStringField(value: v, fieldNumber: 5)
    }()
    case .phone?: try {
      guard case .phone(let v)? = self.contactMethod else { preconditionFailure() }
      try visitor.visitSingularStringField(value: v, fieldNumber: 6)
    }()
    case .didcommInvitation?: try {
      guard case .didcommInvitation(let v)? = self.contactMethod else { preconditionFailure() }
      try visitor.visitSingularMessageField(value: v, fieldNumber: 7)
    }()
    case nil: break
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Trinsic_Services_InviteRequest, rhs: Trinsic_Services_InviteRequest) -> Bool {
    if lhs.participant != rhs.participant {return false}
    if lhs.description_p != rhs.description_p {return false}
    if lhs.contactMethod != rhs.contactMethod {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Trinsic_Services_InviteRequest.DidCommInvitation: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = Trinsic_Services_InviteRequest.protoMessageName + ".DidCommInvitation"
  public static let _protobuf_nameMap = SwiftProtobuf._NameMap()

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let _ = try decoder.nextFieldNumber() {
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Trinsic_Services_InviteRequest.DidCommInvitation, rhs: Trinsic_Services_InviteRequest.DidCommInvitation) -> Bool {
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Trinsic_Services_InviteResponse: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".InviteResponse"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "status"),
    10: .standard(proto: "invitation_id"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularEnumField(value: &self.status) }()
      case 10: try { try decoder.decodeSingularStringField(value: &self.invitationID) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if self.status != .success {
      try visitor.visitSingularEnumField(value: self.status, fieldNumber: 1)
    }
    if !self.invitationID.isEmpty {
      try visitor.visitSingularStringField(value: self.invitationID, fieldNumber: 10)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Trinsic_Services_InviteResponse, rhs: Trinsic_Services_InviteResponse) -> Bool {
    if lhs.status != rhs.status {return false}
    if lhs.invitationID != rhs.invitationID {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Trinsic_Services_InvitationStatusRequest: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".InvitationStatusRequest"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .standard(proto: "invitation_id"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularStringField(value: &self.invitationID) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if !self.invitationID.isEmpty {
      try visitor.visitSingularStringField(value: self.invitationID, fieldNumber: 1)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Trinsic_Services_InvitationStatusRequest, rhs: Trinsic_Services_InvitationStatusRequest) -> Bool {
    if lhs.invitationID != rhs.invitationID {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Trinsic_Services_InvitationStatusResponse: SwiftProtobuf.Message, SwiftProtobuf._MessageImplementationBase, SwiftProtobuf._ProtoNameProviding {
  public static let protoMessageName: String = _protobuf_package + ".InvitationStatusResponse"
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    1: .same(proto: "status"),
    2: .standard(proto: "status_details"),
  ]

  public mutating func decodeMessage<D: SwiftProtobuf.Decoder>(decoder: inout D) throws {
    while let fieldNumber = try decoder.nextFieldNumber() {
      // The use of inline closures is to circumvent an issue where the compiler
      // allocates stack space for every case branch when no optimizations are
      // enabled. https://github.com/apple/swift-protobuf/issues/1034
      switch fieldNumber {
      case 1: try { try decoder.decodeSingularEnumField(value: &self.status) }()
      case 2: try { try decoder.decodeSingularStringField(value: &self.statusDetails) }()
      default: break
      }
    }
  }

  public func traverse<V: SwiftProtobuf.Visitor>(visitor: inout V) throws {
    if self.status != .error {
      try visitor.visitSingularEnumField(value: self.status, fieldNumber: 1)
    }
    if !self.statusDetails.isEmpty {
      try visitor.visitSingularStringField(value: self.statusDetails, fieldNumber: 2)
    }
    try unknownFields.traverse(visitor: &visitor)
  }

  public static func ==(lhs: Trinsic_Services_InvitationStatusResponse, rhs: Trinsic_Services_InvitationStatusResponse) -> Bool {
    if lhs.status != rhs.status {return false}
    if lhs.statusDetails != rhs.statusDetails {return false}
    if lhs.unknownFields != rhs.unknownFields {return false}
    return true
  }
}

extension Trinsic_Services_InvitationStatusResponse.Status: SwiftProtobuf._ProtoNameProviding {
  public static let _protobuf_nameMap: SwiftProtobuf._NameMap = [
    0: .same(proto: "Error"),
    1: .same(proto: "InvitationSent"),
    2: .same(proto: "Completed"),
  ]
}