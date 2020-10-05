<template>
  <div class="survey-form">
    <div v-if="errorMessage" class="alert alert-danger">{{ errorMessage }}</div>
    <div class="form-group">
      <label class="col-form-label col-form-label-lg" for="name">Name</label>
      <input class="form-control form-control-lg" type="text" placeholder="Provide the name of the survey..." id="name" v-model="name">
    </div>
    <div class="form-group">
      <label class="col-form-label col-form-label-lg" for="name">Questions</label>
      <input
        v-for="(question, index) in questions"
        :key="index"
        class="form-control form-control-lg mb-3"
        type="text"
        placeholder="Enter a question..."
        :id="'question' + (index + 1)"
        v-model="questions[index]"
      >
      <small class="form-text text-muted">You must provide at least two questions.</small>
      <button class="btn btn-info btn-sm mt-3" @click="addQuestion">Add another question</button>
    </div>
    <button type="button" class="btn btn-success btn-lg mt-3" @click="createSurvey">Create survey</button>
    <router-link :to="{ name: 'Home' }" tag="button" class="btn btn-link btn-sm mt-3 text-danger">Cancel</router-link>
  </div>
</template>

<script>
export default {
  name: 'Survey',
  data() {
    return {
      name: "",
      questions: ["", ""],
      errorMessage: "",
    };
  },
  methods: {
    addQuestion() {
      this.questions.push("");
    },
    createSurvey() {
      this.errorMessage = "";

      if (this.name == "") {
        this.errorMessage = "The name is required.";
        return;
      }

      let questions = [];
      for (let i in this.questions) {
        if (this.questions[i] != "") {
          questions.push({text: this.questions[i]});
        }
      }
      this.questions.filter(el => {
        return el != "";
      });

      if (questions.length < 2) {
        this.errorMessage = "There must be at least two questions.";
        return;
      }

      let survey = {
        name: this.name,
        questions: questions,
      }

      // TODO: Make this URL configurable
      fetch('http://localhost:8081/surveys', {
        method: "POST",
        body: JSON.stringify(survey),
        headers: {"Content-type": "application/json"}
      })
        .then(res => Promise.all([res.status, res.json()]))
        .then(data => {
          if (data[0] != 201) {
            throw new Error("Survey creation call resulted in result code: " + data[0]);
          }
          this.$router.push({ name: 'Survey', params: { id: data[1].id }});
        })
        .catch(error => {
          this.errorMessage = "Unable to create survey. Please try again.";
          console.log(error);
        });
    }
  }
}
</script>
