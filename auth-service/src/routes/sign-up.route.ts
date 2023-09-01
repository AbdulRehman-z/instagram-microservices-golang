import express, { Request, Response, NextFunction } from "express";
import { prisma } from "../services/prisma.service";
import { BadRequestError } from "@underthehoodjs/commonjs";
import jwt from "jsonwebtoken";
import { Password } from "../services/hashing.service";

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

      // sign jwt with user id and email
      const userJwt = jwt.sign(
        {
          id: user.id,
          email: user.email,
        },

        process.env.JWT_KEY
      );

      // create a session obj with jwt property and attach it to the request obj
      req.session = {
        jwt: userJwt,
      };

      res.status(200).json(user);
    } catch (error) {
      next(error);
    }
  }
);

export { router as signUpRouter };
