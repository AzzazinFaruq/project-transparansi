<template>
  <v-card
      class="dashboard-card py-3"
      rounded="lg"
      >
    <v-card-title class="d-flex justify-space-between " >
          <h1 class="title">{{ name }}</h1>
        </v-card-title>
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
