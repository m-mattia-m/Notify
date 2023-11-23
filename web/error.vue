<script setup lang="ts">
const props = defineProps({
  error: Object
})

const errorStatus = props.error?.statusCode
const errorMessage = props.error?.statusCode == 404 ? "Page not found" : "A unexpected error occurred"
const errorDescription = props.error?.statusCode == 404 ? "Sorry, we couldn’t find the page you’re looking for." : "Sorry, we couldn't handle the process."
const errorPageTitle = props.error?.statusCode == 404 ? "Not found" : "Unexpected error"
const docsUrl = "https://m-mattia-m.github.io/Notify"

const handleError = () => clearError({ redirect: '/' })

useHead({
  title: `Notify | ${errorStatus} - ${errorPageTitle}`
})
</script>

<template>
  <main class="grid min-h-full place-items-center bg-white px-6 py-24 sm:py-32 lg:px-8">
    <div class="text-center">
      <p class="text-base font-semibold text-indigo-600">{{ errorStatus }}</p>
      <h1 class="mt-4 text-3xl font-bold tracking-tight text-gray-900 sm:text-5xl">{{ errorMessage }}</h1>
      <p class="mt-6 text-base leading-7 text-gray-600">{{ errorDescription }}</p>
      <div class="mt-10 flex items-center justify-center gap-x-6">
        <button @click="handleError" class="rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">Go back home</button>
        <NuxtLink :to="docsUrl" class="text-sm font-semibold text-gray-900">Read documentation <span aria-hidden="true">&rarr;</span></NuxtLink>
      </div>
    </div>
  </main>
</template>

<style scoped>

</style>
