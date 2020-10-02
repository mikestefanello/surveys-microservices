<template>
  <div class="home">
    <div v-if="errorMessage" class="alert alert-danger">{{ errorMessage }}</div> 
    <app-survey v-if="survey" :survey="survey"></app-survey>
  </div>
</template>

<script>
  import Survey from '../components/Survey.vue';
  export default {
    name: 'Survey',
    components: {
      'app-survey': Survey,
    },
    data() {
      return {
        survey: false,
        errorMessage: "",
      };
    },
    created() {
      // TODO: Make this URL controlled via env variable
      fetch("http://localhost:8081/surveys/" + this.$route.params.id)
        .then(res => Promise.all([res.status, res.json()]))
        .then(data => {
          if (data[0] == 404) {
            this.errorMessage = "Survey not found.";
            return;
          }
          this.survey = data[1]
          this.fetchedSurveys = true
        })
        .catch(error => {
          this.errorMessage = "Unable to load survey!";
          console.log(error);
        });
    },
  }
</script>
