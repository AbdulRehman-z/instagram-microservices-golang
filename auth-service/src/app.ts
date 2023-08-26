import express from "express";
import { errorHandlerMiddleware } from "@underthehoodjs/commonjs";
import { signUpRouter } from "./routes/sign-up.route";
import { signInRouter } from "./routes/sign-in.route";
import { signOutRouter } from "./routes/sign-out.route";
import { forgotPasswordRouter } from "./routes/forgot-password.route";
import { updatePasswordRouter } from "./routes/update-password.route";

const app = express();
app.use(express.json());

app.use(signUpRouter);
app.use(signInRouter);
app.use(signOutRouter);
app.use(forgotPasswordRouter);
app.use(updatePasswordRouter);

app.use(errorHandlerMiddleware);

export { app };
