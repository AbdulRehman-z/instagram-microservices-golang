import express, { Request, Response, NextFunction } from "express";
import { supabase } from "../services/supabase.service";
import { BadRequestError } from "@underthehoodjs/commonjs";

const router = express.Router();

router.post(
  "/api/users/sign-up",
  async (req: Request, res: Response, next: NextFunction) => {
    const { email, password, username, bio, age, nickname } = req.body;
    try {
      const { data, error } = await supabase.auth.signUp({
        email: email,
        password: password,
        options: {
          data: {
            username: username,
            age: age,
            bio: bio,
            nickname: nickname,
            status: "offline",
          },
        },
      });

      if (error) {
        throw new BadRequestError(error.message);
      }

      res.status(201).send(data);
    } catch (error) {
      next(error);
    }
  }
);

export { router as signUpRouter };
