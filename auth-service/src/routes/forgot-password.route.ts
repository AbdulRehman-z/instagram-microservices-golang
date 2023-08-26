import express, { Request, Response, NextFunction } from "express";
import { BadRequestError } from "@underthehoodjs/commonjs";
const router = express.Router();

router.post(
  "/api/users/forgot-password",
  async (req: Request, res: Response, next: NextFunction) => {
    try {
    } catch (error) {
      next(error);
    }
  }
);

export { router as forgotPasswordRouter };
