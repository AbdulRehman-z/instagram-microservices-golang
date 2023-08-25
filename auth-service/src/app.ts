import express from "express";
import { errorHandlerMiddleware } from "@underthehoodjs/commonjs";
import { signUpRouter } from "./routes/sign-up.route";

const app = express();
app.use(express.json());

app.use(signUpRouter);

app.use(errorHandlerMiddleware);

export { app };
