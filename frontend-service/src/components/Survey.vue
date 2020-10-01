<template>
  <div class="survey">
    <div class="card bg-light mb-3">
      <div class="card-header">{{ survey.name }}</div>
      <div class="card-body">
        <div v-if="errorMessage" class="alert alert-danger">{{ errorMessage }}</div> 
        <div v-if="statusMessage" class="alert alert-success">{{ statusMessage }}</div> 
        <form>
          <div class="form-group">
            <div v-for="question in survey.questions" :key="question.id" class="form-check">
              <label class="form-check-label">
                <input 
                  type="radio" 
                  class="form-check-input" 
                  :name="'questions-' + survey.id" 
                  :id="'question-' + survey.id + '-' + question.id" 
                  :value="question.id" 
                  v-model="selectedQuestion"
                >
                {{ question.text }}
              </label>
            </div>
          </div>
          <button class="btn btn-info" @click.prevent="vote" :disabled="!selectedQuestion">Vote</button>
          <button type="button" class="btn btn-link">View results</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        selectedQuestion: "",
        statusMessage: "",
        errorMessage: "",
      }
    },
    props: {
      survey: Object
    },
    methods: {
      vote() {
        this.errorMessage = "";
        this.statusMessage = "";

        let body = {
          survey: this.survey.id,
          question: this.selectedQuestion,
        }

        // TODO: Make this URL controlled via env variable
        fetch('http://localhost:8082/vote', {
          method: "POST",
          body: JSON.stringify(body),
          headers: {"Content-type": "application/json"}
        })
        .then(res => res.json())
        .then(() => {
          this.statusMessage = "Your vote has been recorded!";
        })
        .catch(error => {
          this.errorMessage = "Vote request failed. Please try again.";
          console.log(error);
        })
      }
    }
  }
</script>