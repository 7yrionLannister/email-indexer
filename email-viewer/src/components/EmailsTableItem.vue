<script setup>
defineProps({
  emailsJson: {
    type: String,
    required: true,
  },
});
</script>

<template>
  <div>
    <h1>Emails</h1>
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
        <tr v-for="item in JSON.parse(emailsJson)" :key="item">
          <td>
            {{ item["Subject"] }}
          </td>
          <td>
            {{ item["From"] }}
          </td>
          <td>
            {{ item["To"] }}
          </td>
          <td>
            Body
          </td>
        </tr>
      </tbody>
    </table>
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
</style>
