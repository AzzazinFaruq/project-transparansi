<template>
  <div class="mx-4">
    <v-row>
      <v-col>
        <v-card class="pa-4">
          <h2 class="text-h5 font-weight-bold mb-5">History Aktivitas</h2>
          <v-table class="no-divider">
            <thead>
              <tr>
                <th>Tanggal</th>
                <th>User</th>
                <th>Aktivitas</th>
                <th>Status</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in paginatedHistory" :key="item.id">
                <td>{{ item.created_at }}</td>
                <td>{{ item.username }}</td>
                <td>{{ item.aktivitas }}</td>
                <td>
                  <v-badge
                    :color="getStatusColor(item.status)"
                    :text-color="item.status === 'Draft' ? 'black' : 'white'"
                    dot
                    inline

                  >
                  </v-badge>
                  {{ item.status }}
                </td>
              </tr>
            </tbody>
          </v-table>

          <!-- Pagination -->
          <v-divider class="my-3 "></v-divider>
        <div class="d-flex justify-end align-center mb-2">
          <p class="mr-3 d-none d-sm-block">Data/Halaman :</p>
          <div class="mr-5 d-none d-sm-block" style="width: 100px;height: 40px;">
            <v-select
            density="compact"
            class="custom-outline"
            v-model="itemsPerPage"
            :items="[5, 10, 25, 50, 100]"
            variant="outlined"
            @input="updateItemsPerPage"
          ></v-select>
          </div>
          <p class="mr-5">{{ currentPage }} dari {{ totalPages }} Halaman</p>
          <v-pagination
            class="custom-pagination"
            :total-visible="0"
            border
            v-model="currentPage"
            :length="totalPages"
            @input="changePage"
          >
          </v-pagination>
        </div>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      loading: false,
      history: [],
      itemsPerPage: 10,
      currentPage: 1,
    }
  },

  computed: {
    totalPages() {
      return Math.ceil(this.history.length / this.itemsPerPage);
    },
    paginatedHistory() {
      const startIndex = (this.currentPage - 1) * this.itemsPerPage;
      const endIndex = startIndex + this.itemsPerPage;
      return this.history.slice(startIndex, endIndex);
    }
  },

  mounted() {
    this.getHistory();
  },

  methods: {
    async getHistory() {
      try {
        this.loading = true;
        const response = await axios.get('/api/index-log');
        this.history = response.data.data;
      } catch (error) {
        console.error('Error fetching history:', error);
      } finally {
        this.loading = false;
      }
    },

    formatDate(date) {
      const d = new Date(date);
      return d.toLocaleDateString('id-ID', {
        day: 'numeric',
        month: 'long',
        year: 'numeric',
        hour: '2-digit',
        minute: '2-digit'
      });
    },

    getStatusColor(status) {
      switch (status) {
        case 'Draft':
          return 'blue';
        case 'Publish':
          return 'green';
        case 'Sudah Ditanggapi':
          return 'green';
        case 'Belum Ditanggapi':
          return 'orange';
        default:
          return 'grey';
      }
    },

    changePage(newPage) {
      this.currentPage = newPage;
    },

    updateItemsPerPage(value) {
      this.itemsPerPage = value;
      this.currentPage = 1; // Reset ke halaman pertama ketika mengubah items per page
    }
  }
}
</script>

<style scoped>
.v-table {
  width: 100%;
}

.v-table td {
  height: 48px;
}


.v-chip {
  font-size: 12px;
}

/* Loading overlay */
.v-card {
  position: relative;
}

.loading-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(255, 255, 255, 0.7);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1;
}
</style>
