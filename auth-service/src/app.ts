import express from "express";
import { errorHandlerMiddleware } from "@underthehoodjs/commonjs";

const app = express();
app.use(express.json());

app.use(errorHandlerMiddleware);

export { app };
