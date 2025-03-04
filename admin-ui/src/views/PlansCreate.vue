<template>
  <div class="pa-4">
    <h1 class="page__title" v-if="!item">Create plan</h1>
    <v-form v-model="isValid" ref="form">
      <v-row>
        <v-col lg="6" cols="12">
          <v-row align="center">
            <v-col cols="3">
              <v-subheader>Plan type</v-subheader>
            </v-col>
            <v-col cols="9">
              <v-select
                label="Type"
                v-model="plan.type"
                :items="types"
                :rules="generalRule"
              />
            </v-col>
          </v-row>

          <v-row align="center">
            <v-col cols="3">
              <v-subheader>Plan title</v-subheader>
            </v-col>
            <v-col cols="9">
              <v-text-field
                label="Title"
                v-model="plan.title"
                :rules="generalRule"
              />
            </v-col>
          </v-row>

          <v-row align="center">
            <v-col cols="3">
              <v-subheader>Plan kind</v-subheader>
            </v-col>
            <v-col cols="9">
              <v-radio-group
                row
                mandatory
                v-model="plan.kind"
              >
                <v-radio
                  v-for="item of kinds"
                  :key="item"
                  :value="item"
                  :label="item.toLowerCase()"
                />
              </v-radio-group>
            </v-col>
          </v-row>

          <v-row align="center">
            <v-col cols="3">
              <v-subheader>Public</v-subheader>
            </v-col>
            <v-col cols="9">
              <v-switch v-model="plan.public" />
            </v-col>
          </v-row>

          <v-divider />

          <v-tabs v-model="form.title" background-color="background-light">
            <v-tab
              draggable="true"
              active-class="background"
              v-for="(title, i) of form.titles"
              :key="title"
              @drag="(e) => dragTab(e, i)"
              @dragstart="dragTabStart"
              @dragend="dragTabEnd"
              @dblclick="edit = {
                isVisible: true,
                title
              }"
            >
              {{ title }}
              <v-icon
                small
                right
                color="error"
                @click="removeConfig(title)"
              >
                mdi-close
              </v-icon>
            </v-tab>
            <v-text-field
              dense
              outlined
              :label="(edit.isVisible)
                ? `Edit ${edit.title}`
                : 'New config'
              "
              class="ml-2 mt-1 mw-20"
              v-if="isVisible || edit.isVisible"
              @change="addConfig"
            />
            <v-icon
              v-else
              class="ml-2"
              @click="isVisible = true"
            >
              mdi-plus
            </v-icon>
          </v-tabs>

          <v-divider />

          <v-subheader v-if="form.titles.length > 0">
            To edit the title, double-click the LMB
          </v-subheader>

          <v-tabs-items v-model="form.title">
            <v-tab-item
              v-for="(title, i) of form.titles"
              :key="title"
            >
              <component
                :is="template"
                :keyForm="title"
                :resource="plan.resources[i]"
                :product="plan.products[title]"
                :preset="preset(i)"
                @change:resource="(data) => changeResource(i, data)"
                @change:product="(data) => changeProduct(title, data)"
              />
            </v-tab-item>
          </v-tabs-items>
        </v-col>
      </v-row>
      
      <v-row>
        <v-col>
          <v-btn
            class="mr-2"
            color="background-light"
            :loading="isLoading"
            :disabled="!isTestSuccess"
            @click="tryToSend"
          >
            {{ item ? 'Edit' : 'Create' }}
          </v-btn>
          <v-btn
            class="mr-2"
            :color="testButtonColor"
            @click="testConfig"
          >
            Test
          </v-btn>
        </v-col>
      </v-row>
    </v-form>

    <v-snackbar
      v-model="snackbar.visibility"
      :timeout="snackbar.timeout"
      :color="snackbar.color"
    >
      {{ snackbar.message }}
      <template v-if="snackbar.route && Object.keys(snackbar.route).length > 0">
        <router-link :to="snackbar.route"> Look up. </router-link>
      </template>

      <template v-slot:action="{ attrs }">
        <v-btn
          :color="snackbar.buttonColor"
          text
          v-bind="attrs"
          @click="snackbar.visibility = false"
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </div>
</template>

<script>
import api from '@/api.js';
import snackbar from '@/mixins/snackbar.js';

export default {
  name: 'plansCreate-view',
  mixins: [snackbar],
  props: ['item'],
  data: () => ({
    types: [],
    kinds: ['DYNAMIC', 'STATIC'],
    plan: {
      title: '',
      type: 'custom',
      kind: '',
      public: false,
      resources: [],
      products: {}
    },
    form: {
      title: '',
      titles: []
    },
    edit: {
      isVisible: false,
      title: ''
    },
    generalRule: [v => !!v || 'This field is required!'],

    isVisible: true,
    isValid: false,
    isLoading: false,
    isTestSuccess: false,
    testButtonColor: 'background-light',
  }),
  methods: {
    changeResource(num, { key, value }) {
      try {
        value = JSON.parse(value, num);
      } catch {
        value;
      }

      if (key === 'date') {
        this.setPeriod(value, num);
        return;
      }
      if (this.plan.resources[num]) {
        this.plan.resources[num][key] = value;
      } else {
        this.plan.resources.push({ [key]: value });
      }
    },
    changeProduct(obj, { key, value }) {
      try {
        value = JSON.parse(value);
      } catch {
        value;
      }

      if (key === 'date') {
        this.setPeriod(value, obj);
        return;
      } else if (key === 'resources') {
        this.plan.resources = value;
        return;
      } else if (key === 'amount') {
        key = 'resources';
      }

      if (this.plan.products[obj]) {
        this.plan.products[obj][key] = value;
      } else {
        this.plan.products[obj] = { [key]: value };
      }
    },
    preset(i) {
      const title = this.form.titles[i - 1];

      if (this.plan.products[title]) {
        return this.plan.products[title].resources;
      }
      if (this.plan.type === 'custom') return;
      return {
        cpu: 1,
        ram: 1024,
        ip_public: 0
      }
    },
    dragTabStart(e) {
      const el = document.createElement('div');

      e.dataTransfer.dropEffect = 'move';
      e.dataTransfer.effectAllowed = 'move';
      e.dataTransfer.setDragImage(el, 0, 0);
    },
    dragTab(e, i) {
      const width = parseInt(getComputedStyle(e.target).width);
      const all = Array.from(e.target.parentElement.children);
      const next = Math.round(e.layerX / width) + i;
      const prev = e.target.getAttribute('data-x');

      e.target.style.cssText = `transform: translateX(${e.layerX}px)`;
      e.target.setAttribute('data-x', `${e.layerX}`);
      all.shift();
      all.pop();

      if (!all[next] || next === i) return;

      all[next].style.transition = '0.3s';
      if (prev < e.layerX) {
        if (e.layerX > width / 2) {
          all[next].style.transform = `translateX(-${width}px)`;
        } else {
          all[next].style.transform = '';
        }
      } else if (prev > e.layerX) {
        if (e.layerX > width / 2) {
          all[next].style.transform = '';
        } else {
          all[next].style.transform = `translateX(${width}px)`;
        }
      }

      const titles = [...this.form.titles];
      const [newTitle] = titles.splice(i, 1);

      titles.splice(next, 0, newTitle);
      localStorage.setItem('titles', JSON.stringify(titles));
    },
    dragTabEnd(e) {
      const all = Array.from(e.target.parentElement.children);
      const titles = localStorage.getItem('titles');
      const wrapper = all.shift();

      all.forEach((el) => el.removeAttribute('style'));
      this.form.titles = JSON.parse(titles);
      localStorage.removeItem('titles');

      setTimeout(() => {
        const left = all.find((el) =>
          el.className.includes('tab--active')
        ).offsetLeft;

        wrapper.style.left = `${left}px`;
      });
    },
    addConfig(title) {
      if (this.edit.isVisible) {
        const i = this.form.titles
          .indexOf(this.edit.title);

        this.form.titles[i] = title;
        this.edit.isVisible = false;

        return;
      }

      this.form.titles.push(title);
      this.isVisible = false;
    },
    removeConfig(title) {
      this.form.titles = this.form.titles
        .filter((el) => el !== title);

      if (this.form.titles.length <= 0) {
        this.isVisible = true;
      }
    },
    tryToSend() {
      if (!this.isValid) {
        this.$refs.form.validate();
        this.testButtonColor = 'background-light';
        this.isTestSuccess = false;

        return;
      }

      this.isLoading = true;
      Object.entries(this.plan.products)
        .forEach(([key, form]) => {
          const num = this.form.titles
            .findIndex((el) => el === key);
          
          form.sorter = num;
        });

      const id = this.$route.params?.planId;
      const request = (this.item)
        ? api.plans.update(id, this.plan)
        : api.plans.create(this.plan);

      request.then(() => {
          this.showSnackbarSuccess({
            message: (this.item)
              ? 'Plan edited successfully'
              : 'Plan created successfully'
          });
        })
        .catch((err) => {
          this.showSnackbarError({ message: err });
        })
        .finally(() => {
          this.isLoading = false;
        });
    },
    testConfig() {
      if (!this.isValid) {
        this.$refs.form.validate();
        this.testButtonColor = 'background-light';
        this.isTestSuccess = false;

        this.showSnackbarError({
          message: 'Validation failed!',
        });

        return;
      }

      this.testButtonColor = 'success';
      this.isTestSuccess = true;
    },
    setPeriod(date, res) {
      const period = this.getTimestamp(date);

      if (this.plan.kind === 'DYNAMIC') {
        this.plan.resources[res].period = period;
        this.plan.products = {};
      } else {
        this.plan.products[res].period = period;
        this.plan.resources = [];
      }
    },
    getTimestamp({ day, month, year, quarter, week, time }) {
      year = +year + 1970;
      month = +month + quarter * 3 + 1;
      day = +day + week * 7 + 1;

      if (`${day}`.length < 2) {
        day = '0' + day;
      }
      if (`${month}`.length < 2) {
        month = '0' + month;
      }

      return Date.parse(
        `${year}-${month}-${day}T${time}Z`
      ) / 1000;
    },
    getItem() {
      this.form.titles = [];
      if (Object.keys(this.item).length > 0) {
        this.plan = this.item;
        this.isVisible = false;

        this.item.resources.forEach((el) => {
          this.form.titles.push(el.key);
        });
        Object.keys(this.item.products).forEach((key) => {
          this.form.titles.push(key);
        });
      }
    }
  },
  created() {
    const types = require.context(
      "@/components/modules/",
      true,
      /serviceProviders\.vue$/
    );
    types.keys().forEach((key) => {
      const matched = key.match(
        /\.\/([A-Za-z0-9-_,\s]*)\/serviceProviders\.vue/i
      );
      if (matched && matched.length > 1) {
        const type = matched[1];
        this.types.push(type);
      }
    });

    if (this.item) this.getItem();
  },
  computed: {
    template() {
      let type;
      switch (this.plan.kind) {
        case 'DYNAMIC':
          type = 'resources';
          break;
        default:
          type = 'products';
      }

      return () => import(`@/components/plans_form_${type}.vue`);
    }
  },
  watch: {
    'plan.type'() {
      switch (this.plan.type) {
        case 'ione':
          if (this.plan.kind === 'STATIC') return;

          this.form.titles = ['cpu', 'ram', 'ip_public'];
          this.isVisible = false;
          break;
        default:
          this.form.titles = [];
          this.isVisible = true;
      }
    },
    item() {
      this.getItem();
    }
  }
}
</script>

<style scoped>
.page__title {
  color: var(--v-primary-base);
  font-weight: 400;
  font-size: 32px;
  font-family: "Quicksand", sans-serif;
  line-height: 1em;
  margin-bottom: 10px;
}

.theme--dark.v-tabs-items {
  background: var(--v-background-base);
}

.mw-20 {
  max-width: 150px;
}
</style>