<script setup lang="ts">
import {OidcUserManager} from "~/service/oidc";

const layout = ref("default")

onMounted(async () => {
  const user = await OidcUserManager(useRuntimeConfig()).getUser()
  if (user !== null && !user.expired) {
    layout.value = "app"
  }
})

</script>

<template>
  <div class="w-full h-full bg-gray-100">
    <NuxtLayout :name="layout">
      <NuxtPage/>
    </NuxtLayout>
  </div>
</template>

<style scoped>

</style>