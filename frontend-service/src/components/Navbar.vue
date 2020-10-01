<template>
  <div class="navbar navbar-expand-lg fixed-top navbar-dark bg-primary">
      <div class="container">
        <router-link :to="{ name: 'Home' }" tag="a" class="navbar-brand font-weight-bold">Surveys</router-link>
        <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarResponsive">
          <ul class="navbar-nav">
            <li class="nav-item">
              <router-link :to="{ name: 'NewSurvey' }" tag="a" class="nav-link">+ Create a new survey</router-link>
            </li>
          </ul>
        </div>
      </div>
    </div>
</template>