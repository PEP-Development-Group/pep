import Vue from 'vue'
import App from './App.vue'
// 引入element
// import ElementUI from 'element-ui';
// import 'element-ui/lib/theme-chalk/index.css';
// // 全局配置elementui的dialog不能通过点击遮罩层关闭
// ElementUI.Dialog.props.closeOnClickModal.default = false

// Vue.use(ElementUI);

import {
    Button,
    Select,
    Aside,
    Autocomplete,
    Alert,
    Container,
    Menu,
    Submenu,
    Row,
    Col,
    Upload,
    MenuItem,
    Card,
    Avatar,
    Form,
    FormItem,
    Divider,
    Tag,
    Input,
    Message,
    MessageBox,
    Notification,
    Dialog,
    Loading,
    Icon,
    Popover,
    Header,
    Table,
    TableColumn,
    Link,
    Dropdown,
    DropdownMenu,
    DropdownItem,
    Breadcrumb,
    BreadcrumbItem,
    Main,
    Scrollbar,
    Radio,
    RadioButton,
    InputNumber,
    Collapse,
    TimePicker,
    TimeSelect,
    DatePicker,
    PageHeader,
    Tabs,
    Popconfirm,
    Tooltip,
    RadioGroup,
    Pagination,
    CollapseItem,
    Progress,
    Checkbox
} from "element-ui";

Vue.use(Progress);
Vue.use(Checkbox);
Vue.use(CollapseItem);
Vue.use(Pagination);
Vue.use(RadioGroup);
Vue.use(Popconfirm);
Vue.use(Tooltip);
Vue.use(Tabs);
Vue.use(PageHeader);
Vue.use(TimePicker)
Vue.use(DatePicker)
Vue.use(TimeSelect)
Vue.use(Container)
Vue.use(Collapse)
Vue.use(InputNumber)
Vue.use(Scrollbar)
Vue.use(BreadcrumbItem);
Vue.use(Breadcrumb);
Vue.use(Main);
Vue.use(Radio);
Vue.use(RadioButton);
Vue.use(Button);
Vue.use(Select);
Vue.use(Avatar);
Vue.use(Aside);
Vue.use(MenuItem);
Vue.use(Card);
Vue.use(FormItem);
Vue.use(Form);
Vue.use(Divider);
Vue.use(Tag);
Vue.use(Input);
Vue.use(Upload);
Vue.use(Menu);
Vue.use(Submenu);
Vue.use(Row);
Vue.use(Col);
Vue.use(Autocomplete);
Vue.use(Alert);
Vue.use(Container);
Vue.use(Dialog);
Vue.use(Icon);
Vue.use(Popover);
Vue.use(Header);
Vue.use(Table);
Vue.use(TableColumn);
Vue.use(Link);
Vue.use(DropdownMenu);
Vue.use(DropdownItem);
Vue.use(Dropdown);
Vue.use(Loading)

Vue.prototype.$loading = Loading.service;
Vue.prototype.$msgbox = MessageBox;
Vue.prototype.$alert = MessageBox.alert;
Vue.prototype.$confirm = MessageBox.confirm;
Vue.prototype.$prompt = MessageBox.prompt;
Vue.prototype.$notify = Notification;
Vue.prototype.$message = Message;
Dialog.props.closeOnClickModal.default = false

// 引入封装的router
import router from '@/router/index'

import '@/permission'
import {store} from '@/store/index'

Vue.config.productionTip = false

// 路由守卫
import Bus from '@/utils/bus.js'

Vue.use(Bus)


import {auth} from '@/directive/auth'
// 按钮权限指令
auth(Vue)

export default new Vue({
    render: h => h(App),
    router,
    store
}).$mount('#app')

import md5 from 'js-md5'

Vue.prototype.$md5 = md5