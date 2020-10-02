<template>
  <div class="survey">
    <div class="card bg-light mb-3">
      <div class="card-header">{{ survey.name }}</div>
      <div class="card-body">
        <div v-if="errorMessage" class="alert alert-danger">{{ errorMessage }}</div> 
        <div v-if="statusMessage" class="alert alert-success">{{ statusMessage }}</div>
        <template v-if="showQuestions">
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
        </template>
        <template v-else>
          <ul class="list-group">
            <li v-for="result in results" :key="result.id" class="list-group-item d-flex justify-content-between align-items-center">
              {{ result.text }}
              <span class="badge badge-success badge-pill">{{ result.votes }}</span>
            </li>
          </ul>
        </template>
      </div>
      <div class="card-footer text-muted">
        <template v-if="showQuestions">
          <button class="btn btn-info" @click.prevent="vote" :disabled="!selectedQuestion">Vote</button>
          <button type="button" class="btn btn-link" @click="viewResults">View results</button>
        </template>
        <template v-else>
          <button type="button" class="btn btn-link" @click="showQuestions = true">Go back</button>
        </template>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        selectedQuestion: "",
        showQuestions: true,
        results: [],
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
          this.selectedQuestion = false
        })
        .catch(error => {
          this.errorMessage = "Vote request failed. Please try again.";
          console.log(error);
        })
      },
      getResults() {
        this.errorMessage = "";
        this.statusMessage = "";

        // TODO: Make this URL controlled via env variable
        fetch('http://localhost:8082/results/' + this.survey.id)
          .then(res => res.json())
          .then(data => {
            this.results = [];
            for (let i in this.survey.questions) {
              this.results[i] = this.survey.questions[i];
              this.results[i].votes = 0;

              for (let n in data.results) {
                if (data.results[n].question == this.survey.questions[i].id) {
                  this.results[i].votes = data.results[n].votes;
                  break;
                }
              }
            }
          })
          .catch(error => {
            this.errorMessage = "Cannot get survey results. Please try again.";
            console.log(error);
          });
      },
      viewResults() {
        this.getResults();
        this.showQuestions = false;
      }
    }
  }
</script>