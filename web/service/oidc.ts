import { UserManager } from 'oidc-client-ts';
import {type RuntimeConfig} from "nuxt/schema";

export const OidcUserManager = (config: RuntimeConfig): UserManager => {
    return new  UserManager({
        authority: `${config.public.oidcIssuer}`,
        client_id:  `${config.public.oidcClientId}`,
        client_secret:  `${config.public.oidcClientId}`,
        redirect_uri:  `${config.public.appUrl}/callback`,
        scope: `openid profile email offline_access`,
        response_type: 'code',
        automaticSilentRenew: true,
        loadUserInfo: true,
    });
}

