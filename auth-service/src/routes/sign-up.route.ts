import express, { Request, Response, NextFunction } from "express";
import { prisma } from "../services/prisma.service";
import { BadRequestError } from "@underthehoodjs/commonjs";
import { Password } from "../services/hashing-service";

const router = express.Router();

router.post(
  "/api/users/sign-up",
  async (req: Request, res: Response, next: NextFunction) => {
    const { email, password, username, bio, age, nickname } = req.body;
    try {
      const userExists = await prisma.user.findFirst({
        where: {
          email: String(email),
        },
      });

      if (userExists) {
        throw new BadRequestError("User already exists");
      }

      const hashedPassword = Password.genPasswordHash(password);

      const user = await prisma.user.create({
        data: {
          email: String(email),
          password: hashedPassword,
        },
      });

      if (!user) {
        throw new BadRequestError("Error registering user");
      }

      res.status(200).json(user);
    } catch (error) {
      next(error);
    }
  }
);

export { router as signUpRouter };
