<template>
  <div class="demo-component">
    <h1>{{ title }}</h1>
    <button @click="incrementCount">Count: {{ count }}</button>
  </div>
</template>

<script>
export default {
  name: "DemoComponent",
  data() {
    return {
      title: "Demo Component",
      count: 0,
    };
  },
  methods: {
    incrementCount() {
      this.count++;
    },
  },
};
</script>

<style scoped>
.demo-component {
  padding: 20px;
  text-align: center;
}

button {
  padding: 8px 16px;
  margin: 10px;
  border-radius: 4px;
  background-color: #4caf50;
  color: white;
  border: none;
  cursor: pointer;
}
</style>
