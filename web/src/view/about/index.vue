<template>
    <div class="container">
        <h1 class="header">Designers and Developers.</h1>
        <el-card class="p">
            <el-avatar :src="imgljk" class="org-img" shape="square"></el-avatar>
            <div class="right">
                <span class="desc">Name: &nbsp;&nbsp;&nbsp; wasabi</span>
                <br />

                <span class="desc">
                    Blog: &nbsp;&nbsp;&nbsp; https://kcode.icu/
                </span>
                <br />

                <span class="desc">
                    Discriptor: &nbsp; &nbsp; &nbsp; He is a gopher!
                </span>
            </div>
        </el-card>

        <el-card class="p"
            ><el-avatar
                :src="imgdjh"
                class="org-img"
                shape="square"
            ></el-avatar>
            <div class="right">
                <span class="desc">Name: &nbsp;&nbsp;&nbsp; BladeHiker</span>
                <br />

                <span class="desc">
                    Blog: &nbsp;&nbsp;&nbsp; https://kcode.icu/
                </span>
                <br />

                <span class="desc">
                    Discriptor: &nbsp; &nbsp; &nbsp; He is a gopher!
                </span>
            </div>
        </el-card>

        <el-card class="p"
            ><el-avatar
                :src="imgwyl"
                class="org-img"
                shape="square"
            ></el-avatar>
            <div class="right">
                <span class="desc">Name: &nbsp;&nbsp;&nbsp; 霜序廿</span>
                <br />

                <span class="desc">
                    Blog: &nbsp;&nbsp;&nbsp; https://kcode.icu/
                </span>
                <br />

                <span class="desc">
                    Discriptor: &nbsp; &nbsp; &nbsp; He is a gopher!
                </span>
            </div>
        </el-card>
    </div>
</template>

<script>
export default {
    name: "About",
    data() {
        return {
            imgwyl: require("@/assets/imgwyl.jpg"),
            imgdjh: require("@/assets/imgdjh.jpg"),
            imgljk: require("@/assets/imgljk.jpg"),
        };
    },
    methods: {},
};
</script>

<style scoped>
.p {
    padding-bottom: 15px;
    margin-bottom: 15px;
}

.container {
    padding: 60px;
}

.desc {
    margin-left: 30px;
    font-size: 20px;
    margin-bottom: 50px;
}

.org-img {
    height: 100px;
    width: 100px;
    float: left;
}

.header {
    font-size: 30px;
    font-weight: 800;
    margin-bottom: 30px;
}

.el-card b {
    font-size: x-large;
    font-weight: lighter;
}

.el-card p {
    font-size: small;
    line-height: normal;
}

.avatar-img {
    float: left;
    height: 40px;
    width: 40px;
    border-radius: 50%;
    -webkit-border-radius: 50%;
    -moz-border-radius: 50%;
    margin-top: 15px;
}

.dom-center {
    margin-left: 50%;
    transform: translateX(-50%);
}
</style>