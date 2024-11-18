<template>
  <v-container>
    <v-row>
      <v-col>
        <h2 class="text-h5 font-weight-bold mb-5">History Aktivitas</h2>

        <v-card>
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
                  <v-chip
                    :color="getStatusColor(item.status)"
                    :text-color="item.status === 'Draft' ? 'black' : 'white'"
                    size="small"
                  >
                    {{ item.status }}
                  </v-chip>
                </td>
              </tr>
            </tbody>
          </v-table>

          <!-- Pagination -->
          <div class="d-flex justify-center align-center pa-4">
            <v-pagination
              v-model="page"
              :length="totalPages"
              :total-visible="7"
              rounded="circle"
            ></v-pagination>
          </div>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      loading: false,
      history: [],
      page: 1,
      itemsPerPage: 10,
    }
  },

  computed: {
    totalPages() {
      return Math.ceil(this.history.length / this.itemsPerPage);
    },
    paginatedHistory() {
      const start = (this.page - 1) * this.itemsPerPage;
      const end = start + this.itemsPerPage;
      return this.history.slice(start, end);
    },
    startNumber() {
      return (this.page - 1) * this.itemsPerPage + 1;
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
          return 'warning';
        case 'Publish':
          return 'success';
        case 'Ditolak':
          return 'error';
        case 'Menunggu':
          return 'info';
        default:
          return 'grey';
      }
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
