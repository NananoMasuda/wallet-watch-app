<template>
  <div>
      <router-view></router-view>
  </div>
</template>

<style lang="scss">
  @import '@/styles/variables.scss';

  body, html {
    background-color: $color-bg;
  }

  body, p, input, select, textarea, button {
    font-family: 'Montserrat', sans-serif;
    letter-spacing: -0.2px;
    font-size: $ruler;
  }

  div, p {
  color: $color-shadow;
  text-shadow: 1px 1px 1px $color-white;
  }

  form {
  padding: $ruler;
  width: $ruler*20;
  margin: 0 auto;
  }

  button, input {
  border: 0;
  outline: 0;
  font-size: $ruler;
  border-radius: $ruler*20;
  padding: $ruler;
  background-color:$color-bg;
  text-shadow: 1px 1px 0 $color-white;
  }

  label {
  display: block;
  margin-bottom: $ruler*1.5;
  width: 100%;
  }

  input {
  margin-right: $ruler/2;
  box-shadow:  inset 2px 2px 5px $color-shadow, inset -5px -5px 10px $color-white;
  width: 100%;
  box-sizing: border-box;
  transition: all 0.2s ease-in-out;
  appearance: none;
  -webkit-appearance: none;

  &:focus {
    box-shadow:  inset 1px 1px 2px $color-shadow, inset -1px -1px 2px $color-white;
  }
  }

  button {
  color:#61677C;
  font-weight: bold;
  box-shadow: -5px -5px 20px $color-white,  5px 5px 20px $color-shadow;
  transition: all 0.2s ease-in-out;
  cursor: pointer;
  font-weight: 600;
  
  &:hover {
    box-shadow: -2px -2px 5px $color-white, 2px 2px 5px $color-shadow;
  }
  
  &:active {
    box-shadow: inset 1px 1px 2px $color-shadow, inset -1px -1px 2px $color-white;
  }
  }

  h1,p {
    text-align: center;
  }
</style>