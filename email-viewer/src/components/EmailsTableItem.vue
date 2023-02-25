<script setup>
import EmailBody from "./EmailBodyItem.vue";
</script>

<template>
  <div id="page-content">
    <div>
      <h1>Emails</h1>
      <input v-model="user" type="text" name="user" id="user" @keyup="updateUser()">
      <button @click="showUserEmails()">Search</button>
      <table>
        <thead>
          <tr>
            <th scope="col">
              Subject
            </th>
            <th scope="col">
              From
            </th>
            <th scope="col">
              To
            </th>
            <th scope="col">
              View body
            </th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="item in emails" :key="item">
            <td>
              {{ item["_source"]["Subject"] }}
            </td>
            <td>
              {{ item["_source"]["From"] }}
            </td>
            <td>
              {{ item["_source"]["To"] }}
            </td>
            <td>
              <a @click="showBody(item['_source']);">Body</a>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <div id="email-body">
      <EmailBody :key="currentBody" :body="currentBody" :subject="currentSubject" />
    </div>
  </div>
</template>

<style scoped>
h1 {
  font-weight: 500;
  font-size: 2.6rem;
  top: -10px;
}

h3 {
  font-size: 1.2rem;
}

table {
  border-collapse: collapse;
}

th,
td {
  border: 0.1rem solid cyan;
  padding: 0.5rem;
}

th {
  font-style: oblique;
  font-weight: bold;
  font-size: larger;
  background-color: darkslategray;
}

.greetings h1,
.greetings h3 {
  text-align: center;
}

@media (min-width: 1024px) {

  .greetings h1,
  .greetings h3 {
    text-align: left;
  }
}

#page-content {
  display: flex;
  flex-direction: horizontal;
}
</style>

<script>
import axios from 'axios';
export default {
  data() {
    return {
      emails: [],
      currentUser: "",
      currentBody: "",
      currentSubject: "",
    };
  },

  methods: {
    async getData() {
      try {
        const response = await axios.get(
          "http://localhost:6060/get-emails?name=" + this.user,
          { crossOriginIsolated: false }
        );
        this.emails = response.data["hits"]["hits"];
        console.log(response.data["hits"]["hits"])
      } catch (error) {
        console.log(error);
      }
    },
    showBody(item) {
      this.currentBody = item['body']
      this.currentSubject = item['Subject']
    },
    updateUser() {
      this.currentUser = this.user
      console.log(this.currentUser);
    },
    showUserEmails() {
      console.log(this.currentUser);
      this.getData()
    }
  },

  created() {
    this.getData();
  },
};
</script>