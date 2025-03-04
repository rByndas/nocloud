<template>
  <div>
    <!-- Secrets -->
    <v-card-title class="px-0 mb-3"> Secrets:</v-card-title>
    <v-row v-if="!editing">
      <v-col>
        <v-text-field
          readonly
          :value="template.secrets.group"
          label="group"
          style="display: inline-block; width: 330px"
        >
        </v-text-field>
      </v-col>
      <v-col>
        <v-text-field
          readonly
          :value="template.secrets.host"
          label="host"
          style="display: inline-block; width: 330px"
        >
        </v-text-field>
      </v-col>
      <v-col>
        <v-text-field
          readonly
          :value="template.secrets.user"
          label="user"
          style="display: inline-block; width: 330px"
        >
        </v-text-field>
      </v-col>
      <v-col>
        <v-text-field
          readonly
          :value="template.secrets.pass"
          label="password"
          style="display: inline-block; width: 330px"
          :type="showPassword ? 'text' : 'password'"
          :append-icon="showPassword ? 'mdi-eye-off' : 'mdi-eye'"
          @click:append="() => (showPassword = !showPassword)"
        >
        </v-text-field>
      </v-col>
    </v-row>
    <slot></slot>
    <!-- Hosts -->
    <v-card-title class="px-0 mb-3">Hosts:</v-card-title>
    <v-row class="mb-7">
      <v-col v-if="template.state.meta.hosts.error" cols="12" lg="6">
        <v-alert type="error"> {{ template.state.meta.hosts.error }}</v-alert>
      </v-col>
      <v-col
        v-for="(host, idx) in template.state.meta.hosts"
        :key="idx"
        cols="12"
        v-else
      >
        <v-row>
          <v-col v-if="host.error" cols="12" lg="6">
            <v-alert type="error"> {{ host.error }}</v-alert>
          </v-col>
          <v-col class="order-2 order-lg-1" cols="12" lg="6" v-else>
            <v-row>
              <v-col cols="12">
                <div class="title_progress">
                  <span>CPU</span>
                  <div>
                    <span>{{ host.total_cpu - host.free_cpu }}</span> /
                    <span>{{ host.total_cpu }}</span>
                  </div>
                </div>
                <v-progress-linear
                  :value="
                    Math.round(
                      ((host.total_cpu - host.free_cpu) / host.total_cpu) * 100
                    )
                  "
                  color="green"
                  height="20"
                >
                  <template v-slot:default="{ value }">
                    <strong>{{ value }}%</strong>
                  </template>
                </v-progress-linear>
              </v-col>
              <v-col cols="12">
                <div class="title_progress">
                  <span>Memory</span>
                  <div>
                    <span
                      >{{
                        ((host.total_ram - host.free_ram) / 1048576).toFixed(2)
                      }}
                      GiB</span
                    >
                    /
                    <span>{{ (host.total_ram / 1048576).toFixed(2) }} GiB</span>
                  </div>
                </div>
                <v-progress-linear
                  :value="
                    Math.round(
                      ((host.total_ram - host.free_ram) / host.total_ram) * 100
                    )
                  "
                  color="green"
                  height="20"
                >
                  <template v-slot:default="{ value }">
                    <strong>{{ value }}%</strong>
                  </template>
                </v-progress-linear>
              </v-col>
            </v-row>
          </v-col>
          <v-col cols="12" lg="6" class="order-1 order-lg-2">
            <p>name: {{ host.name }}</p>
            <p>state: {{ host.state }}</p>
            <p>vm_mad: {{ host.vm_mad }}</p>
            <p>im_mad: {{ host.im_mad }}</p>
          </v-col>
        </v-row>
      </v-col>
    </v-row>
    <!-- Networking -->
    <v-card-title class="px-0 mb-3">Networking:</v-card-title>
    <v-row class="mb-7">
      <v-col>
        <v-card-subtitle class="px-0">Public</v-card-subtitle>
        <v-row>
          <v-col
            v-if="template.state.meta.networking.public_vnet.error"
            cols="12"
            lg="6"
          >
            <v-alert type="error">
              {{ template.state.meta.networking.public_vnet.error }}</v-alert
            >
          </v-col>

          <v-col v-else cols="12" lg="6">
            <div class="title_progress">
              <span>
                {{ template.state.meta.networking.public_vnet.name }}
              </span>
              <div>
                <span>{{
                  template.state.meta.networking.public_vnet.used
                }}</span>
                /
                <span>{{
                  template.state.meta.networking.public_vnet.total
                }}</span>
              </div>
            </div>
            <v-progress-linear
              :value="
                Math.round(
                  (template.state.meta.networking.public_vnet.used /
                    template.state.meta.networking.public_vnet.total) *
                    100
                )
              "
              color="green"
              height="20"
              class="mb-10"
            >
              <template v-slot:default="{ value }">
                <strong>{{ value }}%</strong>
              </template>
            </v-progress-linear>
          </v-col>
        </v-row>
        <v-card-subtitle class="px-0">Private</v-card-subtitle>
        <v-row ref="private">
          <v-col
            v-if="template.state.meta.networking.private_vnet.error"
            cols="12"
            lg="6"
          >
            <v-alert type="error">
              {{ template.state.meta.networking.private_vnet.error }}</v-alert
            >
          </v-col>
          <v-col v-else cols="12" lg="6">
            <div class="title_progress">
              <span>
                {{ template.state.meta.networking.private_vnet.name }}
              </span>
              <div>
                <span>{{
                  template.state.meta.networking.private_vnet.used
                }}</span>
                /
                <span>{{
                  template.state.meta.networking.private_vnet.total
                }}</span>
              </div>
            </div>
            <v-progress-linear
              :value="
                Math.round(
                  (template.state.meta.networking.private_vnet.used /
                    template.state.meta.networking.private_vnet.total) *
                    100
                )
              "
              color="green"
              height="20"
              class="mb-10"
            >
              <template v-slot:default="{ value }">
                <strong>{{ value }}%</strong>
              </template>
            </v-progress-linear>
          </v-col>
        </v-row>
      </v-col>
    </v-row>

    <!-- Datastores -->
    <v-card-title class="px-0 mb-3">Datastores:</v-card-title>
    <v-row class="mb-7">
      <v-col v-if="template.state.meta.datastores.error" cols="12" lg="6">
        <v-alert type="error">
          {{ template.state.meta.datastores.error }}</v-alert
        >
      </v-col>
      <v-col
        v-for="(datastor, idx) in template.state.meta.datastores"
        :key="idx"
        v-else
        cols="12"
      >
        <v-row>
          <v-col class="order-2 order-lg-1" cols="12" lg="6" order-2>
            <div class="title_progress">
              <span>{{ datastor.drive_type }}</span>
              <div>
                <span>{{ (datastor.used / 1024).toFixed(2) }}</span> /
                <span>{{ (datastor.total / 1024).toFixed(2) }} GiB</span>
              </div>
            </div>
            <v-progress-linear
              :value="Math.round((datastor.used / datastor.total) * 100)"
              color="green"
              height="20"
            >
              <template v-slot:default="{ value }">
                <strong>{{ value }}%</strong>
              </template>
            </v-progress-linear>
          </v-col>
          <v-col class="order-1 order-lg-2" cols="12" lg="6" order-1>
            <p>name: {{ datastor.name }}</p>
            <p>tm_mad: {{ datastor.tm_mad }}</p>
            <p>ds_mad: {{ datastor.tm_mad }}</p>
          </v-col>
        </v-row>
      </v-col>
      <v-col>
        <p>OS's:</p>
        <div class="newCloud__template">
          <div
            v-for="OS in template.publicData.templates"
            class="newCloud__template-item"
            :key="OS.name"
          >
            <div class="newCloud__template-image">
              <img :src="'img/OS/' + OS.name.replace(/[^a-zA-Z]+/g, '').toLowerCase() + '.png'" :alt="OS.name" />
            </div>
            <div class="newCloud__template-name">
              {{ OS.desc }}
            </div>
          </div>
        </div>
      </v-col>
      <v-col cols="12">
        <p>Vlans:</p>
        <v-tooltip bottom color="info" v-for="(vlan, i) of vlans" :key="i">
          <template v-slot:activator="{ on, attrs }">
            <span
              class="ceil"
              v-bind="attrs"
              v-on="on"
              :class="vlan === 0 ? 'occupied' : 'free'"
            />
          </template>
          <span>{{ i }}</span>
        </v-tooltip>
        <div class="mt-2">
          <v-btn class="mr-2" v-if="counter > 1" @click="counter--">
            less
          </v-btn>
          <v-btn v-if="counter < 8" @click="counter++"> more </v-btn>
        </div>
      </v-col>
    </v-row>
  </div>
</template>

<script>
export default {
  name: "service-provider-ione",
  props: {
    template: { type: Object, required: true },
    editing: { type: Boolean, required: true },
  },
  data: () => ({
    showPassword: false,
    counter: 1,
  }),
  methods: {
    changeWidth() {
      const { clientWidth } = this.$refs.private;
      let cols = 64;

      for (let i = 4; i > 0; i--) {
        if (clientWidth / 30 >= cols) {
          this.$refs.private.style.width = `${cols * 30 + 24}px`;
          break;
        } else {
          cols /= 2;
        }
      }
    },
  },
  mounted() {
    this.changeWidth();
  },
  computed: {
    vlans() {
      const { free_vlans } = this.template?.state.meta.networking.private_vnet;
      let vlans = 0;

      Object.values(free_vlans || {}).forEach((value) => {
        vlans += +value;
      });

      const res = Array.from({ length: 504 * this.counter })
        .fill(1, 0, vlans)
        .fill(0, vlans);

      return res;
    },
  },
  watch: {
    counter() {
      this.changeWidth();
    },
  },
};
</script>
<style lang="scss" scoped>
.newCloud__template {
  display: flex;
  flex-wrap: wrap;
  &.one-line {
    flex-wrap: nowrap;
    justify-content: space-between;
  }
}

.newCloud__template-item {
  width: 150px;
  margin-left: 20px;
  margin-bottom: 10px;
  box-shadow: 3px 2px 6px rgba(189, 188, 188, 0.08),
    0px 0px 8px rgba(65, 64, 64, 0.05);
  border-radius: 15px;
  border: solid 1px white;
  transition: box-shadow 0.2s ease, transform 0.2s ease;
  cursor: pointer;
  text-align: center;
  overflow: hidden;
  display: grid;
  grid-template-columns: 1fr;
  grid-template-rows: max-content auto;
  &:not(:last-child) {
    margin-right: 10px;
  }
   &:first-child {
    margin-left: 0px;
  }
  &:hover {
    box-shadow: 5px 8px 10px rgba(0, 0, 0, 0.08),
      0px 0px 12px rgba(0, 0, 0, 0.05);
  }
}

  .newCloud__template-image {
    padding: 10px;
  }
  .newCloud__template-image img {
    margin: auto;
    width: 90%;
    height: 100%;
  }

.newCloud__template-name {
  padding: 10px;
}

</style>
