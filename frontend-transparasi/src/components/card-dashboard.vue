<template>
  <v-card
      class="dashboard-card py-3"
      rounded="lg"
      >
        <v-card-title class="d-flex justify-space-between">
          <h1 class="title">{{ name }}</h1>
          <a href="/#" class="dots">
            <v-icon class="icon">mdi-dots-vertical</v-icon>
          </a>
        </v-card-title>
        <div class="d-flex mx-4 mb-2 align-center justify-space-between">
          <div :class="classes">
            <div class="icon ">
              <v-icon class="">{{ icon }}</v-icon>
            </div>
          </div>
          <h1 class="nominal">Rp. 126,5 JT</h1>
          <p class="persentase">+35,78%</p>
        </div>
        <div class="d-flex justify-end mr-4 mb-2 caption">Terbaru Oktober 2024</div>
      </v-card>
</template>
<script>
export default {
  props:{
    classes:{
      type : String,
      Required : true
    },
    icon:{
      type : String,
      Required : true
    },
    name:{
      type : String,
      Required : true
    }
  }
}
</script>
