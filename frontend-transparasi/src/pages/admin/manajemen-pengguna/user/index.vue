<template>
  <div class="mt-3">
        <div class="d-flex align-center justify-space-between mr-3 mt-2">
          <v-card-title><b>Aktifitas Terbaru</b></v-card-title>
          <a href="">
            <v-icon class="">mdi-dots-vertical</v-icon>
          </a>
        </div>
        <v-divider class="mx-2"></v-divider>
        <div class="">
          <v-table class="no-divider">
            <thead>
              <tr class="">
                <th class=" font-weight-bold">
                  Nama
                </th>
                <th class=" font-weight-bold">
                  No Hp
                </th>
                <th class=" font-weight-bold">
                  Email
                </th>
                <th class=" font-weight-bold">
                  Action
                </th>
              </tr>
            </thead>
              <tbody>
                <tr
                v-for="item in Userlist"
                :key="item.name"
                >
              <td>{{ item.username}}</td>
              <td>{{ item.no_hp }}</td>
              <td>{{ item.email }}</td>
              <td>
                <v-btn class="rounded-lg" style="width: 35px;height:35px;background-color:#3884B0;color: white;" icon>
                  <v-icon>mdi-pencil</v-icon>
                </v-btn>
                <v-btn class="rounded-lg ml-2" style="width: 35px;height:35px;background-color: #BF3232;color: white;" icon>
                  <v-icon>mdi-delete</v-icon>
                </v-btn>
              </td>

              </tr>
              </tbody>
          </v-table>
        </div>
    </div>
</template>
<script>
import axios from 'axios';

export default {
  data() {
    return {
      headers: [
        { title: "Nama", value: "username" },
        { title: "Jabatan", value: "" },
        { title: "Email", value: "email" },
        { title: "No Hp", value: "no_hp" },
        { title: "Action", value: "actions" },
      ],
      Userlist:[],
    }
  },
  mounted() {
    this.user();
  },
  methods: {
    user(){
      axios.get("/api/user/byrole/3")
      .then(res=>{
        console.log(res.data.data)
        this.Userlist=res.data.data
      })
    }
  },
}
</script>

