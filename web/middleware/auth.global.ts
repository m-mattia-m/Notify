import {OidcUserManager} from "~/service/oidc";

export default defineNuxtRouteMiddleware(async (to, from) => {
    if (process.server) {
        console.log("middleware on server side")
        return
    }


    const runtimeConfig = useRuntimeConfig()
    const oidcUserManager = OidcUserManager(runtimeConfig)
    const user = await oidcUserManager.getUser()

    if (to.path.startsWith("/login")){
        console.log("before login")
        await oidcUserManager.signinRedirect()
        console.log("after login")
        return navigateTo('')
    }

    if (to.path.startsWith("/logout") || user?.expired){
        console.log("before logout")
        await oidcUserManager.signoutSilent()
        console.log("after logout")
        return navigateTo('')
    }
    if (to.path.startsWith("/callback")){
        console.log("before callback")
        await oidcUserManager.signinRedirectCallback()
        console.log("after callback")
        return navigateTo('')
    }

    console.log("end of middleware")

})