<script setup lang="ts">
import {ClipboardIcon, ArrowsRightLeftIcon, BookOpenIcon} from '@heroicons/vue/24/outline'
import {OidcUserManager} from "~/service/oidc";

const config = useRuntimeConfig()
const user = await OidcUserManager(config).getUser()
const documentationUrl = "https://m-mattia-m.github.io/Notify"

function copyBearerToClipboard() {
  navigator.clipboard.writeText(`Bearer ${user?.id_token}`)
}

useHead({
  title: `Notify | Dashboard`
})
</script>

<template>
  <div class="flex flex-col md:flex-row">

    <NuxtLink :to="config.public.apiUrl" target="_blank"
            class="mt-2 md:mt-0 rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600 flex flex-row active:bg-indigo-800">
      <ArrowsRightLeftIcon class="h-6 w-6 shrink-0" aria-hidden="true"></ArrowsRightLeftIcon>
      <span class="mt-0.5 ml-2">Swagger</span>
    </NuxtLink>
    <NuxtLink :to="documentationUrl" target="_blank"
              class="mt-2 md:mt-0 ml-0 md:ml-2 rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600 flex flex-row active:bg-indigo-800">
      <BookOpenIcon class="h-6 w-6 shrink-0" aria-hidden="true"></BookOpenIcon>
      <span class="mt-0.5 ml-2">Documentation</span>
    </NuxtLink>
    <button @click="copyBearerToClipboard" type="button"
            class="mt-2 md:mt-0 ml-0 md:ml-2 rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600 flex flex-row active:bg-indigo-800">
      <ClipboardIcon class="h-6 w-6 shrink-0" aria-hidden="true"></ClipboardIcon>
      <span class="mt-0.5 ml-2">Copy Bearer</span>
    </button>
  </div>
</template>

<style scoped>

</style>