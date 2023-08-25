import express, { Request, Response, NextFunction } from "express";
import { supabase } from "../services/supabase.service";
import { BadRequestError } from "@underthehoodjs/commonjs";
const router = express.Router();

router.post(
  "/api/users/forgot-password",
  async (req: Request, res: Response, next: NextFunction) => {
    try {
      const { email } = req.body;

      const { data, error } = await supabase.auth.resetPasswordForEmail(email, {
        redirectTo: "http://insta.dev/api/users/update-password",
      });

      if (error) {
        throw new BadRequestError(error.message);
      }
    } catch (error) {
      next(error);
    }
  }
);

export { router as forgotPasswordRouter };
