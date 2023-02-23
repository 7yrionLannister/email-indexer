<script setup>
import WelcomeItem from "./WelcomeItem.vue";
import DocumentationIcon from "./icons/IconDocumentation.vue";
import ToolingIcon from "./icons/IconTooling.vue";
import EcosystemIcon from "./icons/IconEcosystem.vue";
import CommunityIcon from "./icons/IconCommunity.vue";
import SupportIcon from "./icons/IconSupport.vue";
</script>

<template>
  <WelcomeItem>
    <template #icon>
      <DocumentationIcon />
    </template>
    <template #heading>Documentation</template>

    Vue’s
    <a href="https://vuejs.org/" target="_blank" rel="noopener">official documentation</a>
    provides you with all information you need to get started.
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <ToolingIcon />
    </template>
    <template #heading>Tooling</template>

    This project is served and bundled with
    <a href="https://vitejs.dev/guide/features.html" target="_blank" rel="noopener">Vite</a>. The recommended IDE setup is
    <a href="https://code.visualstudio.com/" target="_blank" rel="noopener">VSCode</a>
    +
    <a href="https://github.com/johnsoncodehk/volar" target="_blank" rel="noopener">Volar</a>. If you need to test your
    components and web pages, check out
    <a href="https://www.cypress.io/" target="_blank" rel="noopener">Cypress</a>
    and
    <a href="https://on.cypress.io/component" target="_blank">Cypress Component Testing</a>.

    <br />

    More instructions are available in <code>README.md</code>.
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <EcosystemIcon />
    </template>
    <template #heading>Ecosystem</template>

    Get official tools and libraries for your project:
    <a href="https://pinia.vuejs.org/" target="_blank" rel="noopener">Pinia</a>,
    <a href="https://router.vuejs.org/" target="_blank" rel="noopener">Vue Router</a>,
    <a href="https://test-utils.vuejs.org/" target="_blank" rel="noopener">Vue Test Utils</a>, and
    <a href="https://github.com/vuejs/devtools" target="_blank" rel="noopener">Vue Dev Tools</a>. If you need more
    resources, we suggest paying
    <a href="https://github.com/vuejs/awesome-vue" target="_blank" rel="noopener">Awesome Vue</a>
    a visit.
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <CommunityIcon />
    </template>
    <template #heading>Community</template>

    Got stuck? Ask your question on
    <a href="https://chat.vuejs.org" target="_blank" rel="noopener">Vue Land</a>, our official Discord server, or
    <a href="https://stackoverflow.com/questions/tagged/vue.js" target="_blank" rel="noopener">StackOverflow</a>. You
    should also subscribe to
    <a href="https://news.vuejs.org" target="_blank" rel="noopener">our mailing list</a>
    and follow the official
    <a href="https://twitter.com/vuejs" target="_blank" rel="noopener">@vuejs</a>
    twitter account for latest news in the Vue world.
  </WelcomeItem>

  <WelcomeItem>
    <template #icon>
      <SupportIcon />
    </template>
    <template #heading>Support Vue</template>

    As an independent project, Vue relies on community backing for its
    sustainability. You can help us by
    <a href="https://vuejs.org/sponsor/" target="_blank" rel="noopener">becoming a sponsor</a>.
  </WelcomeItem>

  <div class="flex flex-col">
    <div class="overflow-x-auto">
      <div class="p-1.5 w-full inline-block align-middle">
        <div class="overflow-hidden border rounded-lg">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th scope="col" class="px-6 py-3 text-xs font-bold text-left text-gray-500 uppercase">
                  ID
                </th>
                <th scope="col" class="px-6 py-3 text-xs font-bold text-left text-gray-500 uppercase">
                  Name
                </th>
                <th scope="col" class="px-6 py-3 text-xs font-bold text-left text-gray-500 uppercase">
                  Email
                </th>
                <th scope="col" class="px-6 py-3 text-xs font-bold text-right text-gray-500 uppercase">
                  Edit
                </th>
                <th scope="col" class="px-6 py-3 text-xs font-bold text-right text-gray-500 uppercase">
                  Delete
                </th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-200">
              <tr>
                <td class="px-6 py-4 text-sm font-medium text-gray-800 whitespace-nowrap">
                  1
                </td>
                <td class="px-6 py-4 text-sm text-gray-800 whitespace-nowrap">
                  Jone Doe
                </td>
                <td class="px-6 py-4 text-sm text-gray-800 whitespace-nowrap">
                  jonne62@gmail.com
                </td>
                <td class="px-6 py-4 text-sm font-medium text-right whitespace-nowrap">
                  <a class="text-green-500 hover:text-green-700" href="#">
                    Edit
                  </a>
                </td>
                <td class="px-6 py-4 text-sm font-medium text-right whitespace-nowrap">
                  <a class="text-red-500 hover:text-red-700" href="#">
                    Delete
                  </a>
                </td>
              </tr>
              <tr>
                <td class="px-6 py-4 text-sm font-medium text-gray-800 whitespace-nowrap">
                  2
                </td>
                <td class="px-6 py-4 text-sm text-gray-800 whitespace-nowrap">
                  Jone Doe
                </td>
                <td class="px-6 py-4 text-sm text-gray-800 whitespace-nowrap">
                  jonne62@gmail.com
                </td>
                <td class="px-6 py-4 text-sm font-medium text-right whitespace-nowrap">
                  <a class="text-green-500 hover:text-green-700" href="#">
                    Edit
                  </a>
                </td>
                <td class="px-6 py-4 text-sm font-medium text-right whitespace-nowrap">
                  <a class="text-red-500 hover:text-red-700" href="#">
                    Delete
                  </a>
                </td>
              </tr>
              <tr>
                <td class="px-6 py-4 text-sm font-medium text-gray-800 whitespace-nowrap">
                  3
                </td>
                <td class="px-6 py-4 text-sm text-gray-800 whitespace-nowrap">
                  Jone Doe
                </td>
                <td class="px-6 py-4 text-sm text-gray-800 whitespace-nowrap">
                  jonne62@gmail.com
                </td>
                <td class="px-6 py-4 text-sm font-medium text-right whitespace-nowrap">
                  <a class="text-green-500 hover:text-green-700" href="#">
                    Edit
                  </a>
                </td>
                <td class="px-6 py-4 text-sm font-medium text-right whitespace-nowrap">
                  <a class="text-red-500 hover:text-red-700" href="#">
                    Delete
                  </a>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>
