<template>
  <div class="home">
    <div v-if="errorMessage" class="alert alert-danger">{{ errorMessage }}</div> 
    <div v-if="fetchedSurveys && !surveys.length" class="alert alert-warning">There are currently no surveys.</div>
    <app-survey v-for="survey in surveys" :survey="survey" :key="survey.id"></app-survey>
  </div>
</template>

<script>
  import Survey from '../components/Survey.vue';
  export default {
    name: 'Home',
    components: {
      'app-survey': Survey,
    },
    data() {
      return {
        surveys: [],
        errorMessage: "",
        fetchedSurveys: false,
      };
    },
    created() {
      // TODO: Make this URL controlled via env variable
      fetch("http://localhost:8081/surveys")
      .then(res => res.json())
      .then(data => {
        this.surveys = data
        this.fetchedSurveys = true
      })
      .catch(error => {
        this.errorMessage = "Unable to load surveys!";
        console.log(error);
      });
    },
  }
</script>
