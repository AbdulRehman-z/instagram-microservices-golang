import express, { Request, Response, NextFunction } from "express";
import { BadRequestError, NotFoundError } from "@underthehoodjs/commonjs";
import { prisma } from "../services/prisma.service";

const router = express.Router();

router.post(
  "/api/users/sign-in",
  async (req: Request, res: Response, next: NextFunction) => {
    const { email, password } = req.body;
    try {
      const userExists = await prisma.user.findFirst({
        where: {
          email: String(email),
        },
      });

      if (userExists) {
        throw new BadRequestError("User already exists");
      }
    } catch (error) {
      next(error);
    }
  }
);

export { router as signInRouter };
