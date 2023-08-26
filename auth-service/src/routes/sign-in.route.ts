import express, { Request, Response, NextFunction } from "express";
import { BadRequestError, NotFoundError } from "@underthehoodjs/commonjs";

const router = express.Router();

router.post(
  "/api/users/sign-in",
  async (req: Request, res: Response, next: NextFunction) => {
    const { email, password } = req.body;
    try {
    } catch (error) {
      next(error);
    }
  }
);

export { router as signInRouter };
