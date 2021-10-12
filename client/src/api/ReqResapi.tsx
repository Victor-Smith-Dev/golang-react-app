import axios from "axios";

//TODO: Api request
export const ReqResApi = axios.create({
  baseURL: "http://localhost:3000/api/v1",
  headers: {
    "Access-Control-Allow-Origin": "*",
    "Access-Control-Allow-Methods": "GET,PUT,POST,DELETE,PATCH,OPTIONS",
    "Content-Type": "application/json",
  },
});
