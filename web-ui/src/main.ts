import { createApp } from 'vue'
import App from './App.vue'
import PrimeVue from 'primevue/config';
import router from './router'

import Menubar from 'primevue/menubar';
import Button from 'primevue/button';
import Password from 'primevue/password';
import InputText from 'primevue/inputtext';
import Card from 'primevue/card';
import InlineMessage from 'primevue/inlinemessage';

import 'primevue/resources/themes/saga-blue/theme.css';
import 'primevue/resources/primevue.min.css';
import 'primeicons/primeicons.css';

import 'primeflex/primeflex.css';

createApp(App)
    .use(router)
    .use(PrimeVue)
    .component('Menubar', Menubar)
    .component('Button', Button)
    .component('Password', Password)
    .component('InputText', InputText)
    .component('Card', Card)
    .component('InlineMessage', InlineMessage)
    .mount('#app')
