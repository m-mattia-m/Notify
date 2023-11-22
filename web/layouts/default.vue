<script setup lang="ts">
import {ChevronRightIcon} from '@heroicons/vue/20/solid'

const githubRepositoryUrl = "https://github.com/m-mattia-m/Notify/"
const documentationUrl = "https://m-mattia-m.github.io/Notify"

</script>

<template>
  <div class="bg-white">
    <div class="relative isolate overflow-hidden bg-gradient-to-b from-indigo-100/20">
      <div class="mx-auto max-w-7xl pb-24 pt-10 sm:pb-32 lg:grid lg:grid-cols-2 lg:gap-x-8 lg:px-8 lg:py-32">
        <div class="px-6 lg:px-0 lg:pt-4">
          <div class="mx-auto max-w-2xl">
            <div class="max-w-lg">
              <img class="h-11" src="https://tailwindui.com/img/logos/mark.svg?color=indigo&amp;shade=600"
                   alt="Your Company"/>
              <div class="mt-24 sm:mt-32 lg:mt-16">
                <a href="#" class="inline-flex space-x-6">
                  <NuxtLink :to="githubRepositoryUrl"
                            class="rounded-full bg-indigo-600/10 px-3 py-1 text-sm font-semibold leading-6 text-indigo-600 ring-1 ring-inset ring-indigo-600/10">
                    Open Source
                  </NuxtLink>
                  <span class="inline-flex items-center space-x-2 text-sm font-medium leading-6 text-gray-600">
                    <NuxtLink :to="documentationUrl">Documentation</NuxtLink>
                    <ChevronRightIcon class="h-5 w-5 text-gray-400" aria-hidden="true"/>
                  </span>
                </a>
              </div>
              <h1 class="mt-10 text-4xl font-bold tracking-tight text-gray-900 sm:text-6xl">
                Frontend Messaging Made Simple
              </h1>
              <p class="mt-6 text-lg leading-8 text-gray-600">
                Seamlessly send secure messages to your provider—effortless communication without the hassle of API
                keys.
              </p>
              <div class="mt-10 flex items-center gap-x-6">
                <NuxtLink to="/login"
                          class="rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">
                  Login
                </NuxtLink>
                <NuxtLink :to="githubRepositoryUrl" class="text-sm font-semibold leading-6 text-gray-900">
                  View on GitHub <span aria-hidden="true">→</span>
                </NuxtLink>
              </div>
            </div>
          </div>
        </div>
        <div class="mt-20 sm:mt-24 md:mx-auto md:max-w-2xl lg:mx-0 lg:mt-0 lg:w-screen">
          <div
              class="absolute inset-y-0 right-1/2 -z-10 -mr-10 w-[200%] skew-x-[-30deg] bg-white shadow-xl shadow-indigo-600/10 ring-1 ring-indigo-50 md:-mr-20 lg:-mr-36"
              aria-hidden="true"/>
          <div class="shadow-lg md:rounded-3xl">
            <div class="bg-indigo-500 [clip-path:inset(0)] md:[clip-path:inset(0_round_theme(borderRadius.3xl))]">
              <div
                  class="absolute -inset-y-px left-1/2 -z-10 ml-10 w-[200%] skew-x-[-30deg] bg-indigo-100 opacity-20 ring-1 ring-inset ring-white md:ml-20 lg:ml-36"
                  aria-hidden="true"/>
              <div class="relative px-6 pt-8 sm:pt-16 md:pl-16 md:pr-0">
                <div class="mx-auto max-w-2xl md:mx-0 md:max-w-none">
                  <div class="w-screen overflow-hidden rounded-tl-xl bg-gray-900">
                    <div class="flex bg-gray-800/40 ring-1 ring-white/5">
                      <div class="-mb-px flex text-sm font-medium leading-6 text-gray-400">
                        <div
                            class="border-b border-r border-b-white/20 border-r-white/10 bg-white/5 px-4 py-2 text-white">
                          contact.vue
                        </div>
                        <div class="border-r border-gray-600/10 px-4 py-2">index.vue</div>
                      </div>
                    </div>
                    <div class="px-6 pb-14 pt-6">
                      <p class="tracking-normal font-mono my-1">
                        <span class="text-[#ce8e6d]">const </span>
                        <span class="text-[#57a8f5]">sendNotification </span>
                        <span class="text-white">= </span>
                        <span class="text-[#ce8e6d]">async </span>
                        <span class="text-white">() => </span>
                      </p>
                      <p class="pl-4 tracking-normal font-mono my-1">
                        <span class="text-[#c87dba]">axios</span>
                        <span class="text-white">.</span>
                        <span class="text-[#57a8f5]">post</span>
                        <span class="text-white">(</span>
                      </p>
                      <p class="pl-8 tracking-normal font-mono my-1">
                        <span class="text-[#6aab73]">'https://notify.example.com/notification'</span>
                        <span class="text-white">,</span>
                      </p>
                      <p class="pl-8 tracking-normal font-mono my-1">
                        <span class="text-white">{</span>
                      </p>
                      <p class="pl-12 tracking-normal font-mono my-1">
                        <span class="text-[#c87dba]">"project_id"</span>
                        <span class="text-white">: </span>
                        <span class="text-[#6aab73]">"abc123"</span>
                      </p>
                      <p class="pl-12 tracking-normal font-mono my-1">
                        <span class="text-[#c87dba]">"subject"</span>
                        <span class="text-white">: </span>
                        <span class="text-[#6aab73]">"Contact form"</span>
                      </p>
                      <p class="pl-12 tracking-normal font-mono my-1">
                        <span class="text-[#c87dba]">"message"</span>
                        <span class="text-white">: </span>
                        <span class="text-[#6aab73]">"Hi, I need some technical help."</span>
                      </p>
                      <p class="pl-12 tracking-normal font-mono my-1">
                        <span class="text-[#c87dba]">"target"</span>
                        <span class="text-white">: </span>
                        <span class="text-[#6aab73]">"mailgun:hi@example.com"</span>
                      </p>
                      <p class="pl-8 tracking-normal font-mono my-1">
                        <span class="text-white">}</span>
                      </p>
                      <p class="pl-4 tracking-normal font-mono my-1">
                        <span class="text-white">});</span>
                      </p>
                    </div>
                  </div>
                </div>
                <div class="pointer-events-none absolute inset-0 ring-1 ring-inset ring-black/10 md:rounded-3xl"
                     aria-hidden="true"/>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="absolute inset-x-0 bottom-0 -z-10 h-24 bg-gradient-to-t from-white sm:h-32"/>
    </div>
  </div>
</template>

<style scoped>

</style>