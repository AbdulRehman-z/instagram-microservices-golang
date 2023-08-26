import express, { Request, Response, NextFunction } from "express";
import { BadRequestError } from "@underthehoodjs/commonjs";

const router = express.Router();

router.post(
  "/api/users/sign-up",
  async (req: Request, res: Response, next: NextFunction) => {
    const { email, password, username, bio, age, nickname } = req.body;
    try {
    } catch (error) {
      next(error);
    }
  }
);

export { router as signUpRouter };
