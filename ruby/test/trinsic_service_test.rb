require_relative 'test_helper'
require_relative 'vaccine_demo'
require_relative 'ecosystem_demo'
require_relative 'credential_template_demo'
require_relative 'trust_registry_demo'
require 'json'
require 'okapi'
require 'uri'
require 'trinsic_services'
require 'services/service_base'
require 'services/account_service'
require 'services/credential_service'
require 'services/wallet_service'
require 'services/provider_service'
require 'services/credential_template_service'
require 'services/trust_registry_service'
require 'securerandom'

class TrinsicServiceTest < Minitest::Test

  def before_setup
    Okapi::load_native_library
  end

  def test_servicebase_setprofile
    account_service = Trinsic::AccountService.new(Trinsic::trinsic_server)
    assert_raises Exception do
      account_service.metadata(nil)
    end
    account_profile = account_service.sign_in(nil)
    account_service.profile = account_profile
    metadata = account_service.metadata(Services::Provider::V1::InviteRequest.new)
    assert(metadata != nil, "Valid metadata once profile is set")
  end

  def test_providerservice_inviteparticipant
    nil # this test needs ecosystem support
    # account_service = Trinsic::AccountService.new(nil, Trinsic::trinsic_server)
    # account_profile = account_service.sign_in(nil).profile
    # provider_service = Trinsic::ProviderService.new(account_profile, Trinsic::trinsic_server)
    #
    # invite_request = Services::Provider::V1::InviteRequest.new(:description => "I dunno", :email => "does.not.exist@trinsic.id")
    # invite_response = provider_service.invite_participant(invite_request)
    # assert(invite_response != nil)
    # TODO - Verify invitation status response
  end

  def test_report_information
    # This test is about reporting system information, less of a test, and more of result annotations for diagnosing cpu/environment bugs
    puts("Target server= #{Trinsic::trinsic_server}")
  end

  def test_has_a_version_number
    refute_nil ::Trinsic::VERSION
  end

  def test_trinsic_services_demo
    vaccine_demo_run
  end

  def test_ecosystem_demo
    ecosystem_demo_run
  end

  def test_templates_demo
    credential_template_demo_run
  end

  def test_trustregistry_demo
    trust_registry_demo_run
  end
end
