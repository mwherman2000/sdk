import { AccountService, CreateEcosystemRequest, ProviderService, SignInRequest } from "../src";
import { getTestServerOptions } from "./TestData";

require("dotenv").config();

const options = getTestServerOptions();

describe('Demo: Ecosystem Tests', () => {
    beforeAll(async () => {
        let service = new AccountService(options);
        let authToken = await service.signIn(new SignInRequest());

        options.setAuthToken(authToken);
    });

    it('can create ecosystem', async () => {
        let providerService = new ProviderService(options);
        let actualCreate = await providerService.createEcosystem(
            new CreateEcosystemRequest()
                .setDescription("Test ecosystem from Node")
                .setUri("https://example.com"));

        expect(actualCreate).not.toBeNull();
        expect(actualCreate.getEcosystem()).not.toBeNull();
        expect(actualCreate.getEcosystem()!.getId().startsWith("urn:trinsic:ecosystems:")).toBeTruthy();
    });
});